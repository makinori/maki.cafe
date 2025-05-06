package util

import (
	"fmt"
	"hash/crc32"
	"io/fs"
	"log"
	"net/http"
	"strings"
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
