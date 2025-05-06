package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/jinzhu/copier"
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

var portRegexp = regexp.MustCompile(":[0-9]+$")

func HTTPGetIPAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress != "" {
		ipAddress = strings.Split(ipAddress, ",")[0]
		ipAddress = strings.TrimSpace(ipAddress)
		return portRegexp.ReplaceAllString(ipAddress, "")
	}

	return portRegexp.ReplaceAllString(r.RemoteAddr, "")
}

func HTTPGetFullURL(r *http.Request) string {
	var fullUrl url.URL
	copier.Copy(&fullUrl, r.URL)

	fullUrl.Scheme = r.Header.Get("X-Forwarded-Proto")
	if fullUrl.Scheme == "" {
		fullUrl.Scheme = "http"
	}
	fullUrl.Host = r.Host

	return fullUrl.String()
}

func HTTPPlausibleEvent(r *http.Request) bool {
	// if isDev {
	// 	return false
	// }

	// https://plausible.io/docs/events-api

	body, err := json.Marshal(map[string]string{
		"name":    "pageview",
		"url":     HTTPGetFullURL(r),
		"domain":  "maki.cafe",
		"referer": r.Header.Get("Referer"),
	})

	if err != nil {
		log.Println("failed to marshal json for plausible: " + err.Error())
		return false
	}

	req, err := http.NewRequest(
		"POST", "https://ithelpsme.hotmilk.space/api/event",
		bytes.NewReader(body),
	)

	userAgent := r.Header.Get("User-Agent")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/json")

	ipAddress := HTTPGetIPAddress(r)
	if ipAddress != "" {
		r.Header.Add("X-Forwarded-For", ipAddress)
	}

	log.Println(
		"plausible:\n" +
			"  data: " + string(body) + "\n" +
			"  ip: " + ipAddress + "\n" +
			"  ua: " + userAgent,
	)

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println("failed to send plausible event: " + err.Error())
		return false
	}

	return true
}
