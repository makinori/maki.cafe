package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func HTTPWriteWithEncoding(w http.ResponseWriter, r *http.Request, data []byte) {
	// go tool air proxy wont work if encoding
	if ENV_IS_DEV {
		w.Write(data)
		return
	}

	acceptEncoding := r.Header.Get("Accept-Encoding")

	switch {
	case strings.Contains(acceptEncoding, "zstd"):
		data, err := EncodeZstd(data)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		w.Header().Add("Content-Encoding", "zstd")
		w.Write(data)
		return

	case strings.Contains(acceptEncoding, "br"):
		data, err := EncodeBrotli(data)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		w.Header().Add("Content-Encoding", "br")
		w.Write(data)
		return

	}

	w.Write(data)
}

func HTTPServeOptimized(
	w http.ResponseWriter, r *http.Request, data []byte, filename string, allowCache bool,
) {
	contentType := w.Header().Get("Content-Type")

	if allowCache {
		// unset content type incase etag matches
		w.Header().Del("Content-Type")

		// etag := fmt.Sprintf(`W/"%x"`, crc32.ChecksumIEEE(data))
		etag := fmt.Sprintf(`"%x"`, crc32.ChecksumIEEE(data))

		ifMatch := r.Header.Get("If-Match")
		if ifMatch != "" {
			if !InCommaSeperated(ifMatch, etag) && !InCommaSeperated(ifMatch, "*") {
				w.WriteHeader(http.StatusPreconditionFailed)
				return
			}
		}

		ifNoneMatch := r.Header.Get("If-None-Match")
		if ifNoneMatch != "" {
			if InCommaSeperated(ifNoneMatch, etag) || InCommaSeperated(ifNoneMatch, "*") {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		w.Header().Add("ETag", etag)
	} else {
		w.Header().Add("Cache-Control", "no-store")
	}

	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(filename))
		if contentType == "" {
			contentType = http.DetectContentType(data)
		}
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
			slog.Error("failed to file stat", "err", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make([]byte, stat.Size())

		_, err = file.Read(data)
		if err != nil {
			slog.Error("failed to read file", "err", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		HTTPServeOptimized(w, r, data, filename, true)
	}
}

var portRegexp = regexp.MustCompile(":[0-9]+$")

func HTTPGetIPAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Forwarded-For")
	if ipAddress != "" {
		ipAddress = strings.Split(ipAddress, ",")[0]
		ipAddress = strings.TrimSpace(ipAddress)
	} else {
		ipAddress = r.RemoteAddr // from connection
	}

	ipAddress = portRegexp.ReplaceAllString(ipAddress, "")
	ipAddress = strings.TrimPrefix(ipAddress, "[")
	ipAddress = strings.TrimSuffix(ipAddress, "]")

	return ipAddress
}

func HTTPGetFullURL(r *http.Request) url.URL {
	fullUrl := *r.URL // shallow copy

	fullUrl.Scheme = r.Header.Get("X-Forwarded-Proto")
	if fullUrl.Scheme == "" {
		fullUrl.Scheme = "http"
	}
	fullUrl.Host = r.Host

	return fullUrl
}

var ignorePlausibleUserAgents = []string{
	"gatus", // health checks
	"curl",
}

func HTTPPlausibleEventFromImg(incomingReq *http.Request) bool {
	if ENV_IS_DEV {
		return false
	}

	userAgent := incomingReq.Header.Get("User-Agent")
	lowerUserAgent := strings.ToLower(userAgent)

	for _, ignoreUA := range ignorePlausibleUserAgents {
		if strings.Contains(lowerUserAgent, ignoreUA) {
			return false
		}
	}

	notabot, err := NotabotDecode(incomingReq.URL.RawQuery)
	if err != nil {
		slog.Error("failed to get notabot data", "err", err.Error())
		return false
	}

	// https://plausible.io/docs/events-api

	fullUrl := HTTPGetFullURL(incomingReq)
	fullUrl.Path = notabot.Path
	fullUrl.RawQuery = ""

	body, err := json.Marshal(map[string]string{
		"name":     "pageview",
		"url":      fullUrl.String(),
		"domain":   "maki.cafe",
		"referrer": notabot.Ref,
	})

	if err != nil {
		slog.Error("failed to marshal json for plausible", "err", err.Error())
		return false
	}

	plausibleReq, err := http.NewRequest(
		"POST", "https://ithelpsme.hotmilk.space/api/event",
		bytes.NewReader(body),
	)

	plausibleReq.Header.Add("User-Agent", userAgent)
	plausibleReq.Header.Add("Content-Type", "application/json")

	ipAddress := HTTPGetIPAddress(incomingReq)
	if ipAddress != "" {
		// traefik will rewrite this
		// plausibleReq.Header.Add("X-Forwarded-For", ipAddress)
		//https://github.com/plausible/analytics/blob/master/lib/plausible_web/remote_ip.ex
		plausibleReq.Header.Add("X-Plausible-IP", ipAddress)
	}

	if ENV_PLAUSIBLE_DEBUG {
		// isn't even in the plausible code. docs need to be updated
		plausibleReq.Header.Add("X-Debug-Request", "true")
		slog.Info(
			"plausible:\n" +
				"  data: " + string(body) + "\n" +
				"  ip: " + ipAddress + "\n" +
				"  ua: " + userAgent,
		)
	}

	client := http.Client{}
	res, err := client.Do(plausibleReq)
	if err != nil {
		slog.Error("failed to send plausible event", "err", err.Error())
		return false
	}

	if ENV_PLAUSIBLE_DEBUG {
		resData, err := io.ReadAll(res.Body)
		if err != nil {
			slog.Error("failed to read plausible res body", "err", err.Error())
		}
		slog.Info("res: " + string(resData))
	}

	return true
}
