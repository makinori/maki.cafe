package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	_, isDev = os.LookupEnv("DEV")
)

func HTTPWriteWithEncoding(w http.ResponseWriter, r *http.Request, data []byte) {
	acceptEncoding := r.Header.Get("Accept-Encoding")

	switch {
	case strings.Contains(acceptEncoding, "zstd"):
		data, ok := EncodeZstd(data)
		if ok {
			w.Header().Add("Content-Encoding", "zstd")
			w.Write(data)
			return
		}
	case strings.Contains(acceptEncoding, "br"):
		data, ok := EncodeBrotli(data)
		if ok {
			w.Header().Add("Content-Encoding", "br")
			w.Write(data)
			return
		}
	}

	w.Write(data)
}

func HTTPServeOptimized(w http.ResponseWriter, r *http.Request, data []byte) {
	contentType := w.Header().Get("Content-Type")
	fmt.Println(contentType)

	// unset content type incase etag matches
	w.Header().Del("Content-Type")

	// etag := fmt.Sprintf(`W/"%x"`, crc32.ChecksumIEEE(data))
	etag := fmt.Sprintf(`"%x"`, crc32.ChecksumIEEE(data))

	if strings.Contains(r.Header.Get("If-None-Match"), etag) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Add("ETag", etag)

	if contentType == "" {
		contentType = http.DetectContentType(data)
	}

	w.Header().Add("Content-Type", contentType)

	HTTPWriteWithEncoding(w, r, data)
}

func HTTPFileServerOptimized(fs fs.FS) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := r.PathValue("file")

		file, err := fs.Open(filename)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		stat, err := file.Stat()
		if err != nil {
			log.Println("failed to file stat: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make([]byte, stat.Size())

		_, err = file.Read(data)
		if err != nil {
			log.Println("failed to read file: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		HTTPServeOptimized(w, r, data)
	}
}

func HTTPPlausibleEvent(r *http.Request) bool {
	if isDev {
		return false
	}

	// https://plausible.io/docs/events-api

	body, err := json.Marshal(map[string]string{
		"name":     "pageview",
		"url":      r.URL.String(),
		"domain":   "maki.cafe",
		"referrer": r.Header.Get("Referrer"),
	})

	if err != nil {
		log.Println("failed to marshal json for plausible: " + err.Error())
		return false
	}

	log.Println("plausible: " + string(body))

	req, err := http.NewRequest(
		"POST", "https://ithelpsme.hotmilk.space/api/event",
		bytes.NewReader(body),
	)

	req.Header.Add("User-Agent", r.Header.Get("User-Agent"))
	req.Header.Add("Content-Type", "application/json")

	// we're reverse proxying
	ipAddress := req.Header.Get("X-Forwarded-For")
	if ipAddress == "" {
		ipAddress = req.Header.Get("X-Real-IP")
	}
	if ipAddress != "" {
		req.Header.Add("X-Forwarded-For", ipAddress)
	}

	log.Println("ip: " + ipAddress)

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println("failed to send plausible event: " + err.Error())
		return false
	}

	return true
}
