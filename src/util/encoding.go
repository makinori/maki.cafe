package util

import (
	"bytes"
	"fmt"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

func EncodeZstd(input []byte) ([]byte, bool) {
	data := bytes.NewBuffer(nil)

	enc, err := zstd.NewWriter(data)
	if err != nil {
		fmt.Println("failed to init zstd encoder: " + err.Error())
		return []byte{}, false
	}

	_, err = enc.Write(input)
	if err != nil {
		fmt.Println("failed to compress zstd")
		return []byte{}, false
	}

	err = enc.Close()
	if err != nil {
		fmt.Println("failed to close zstd encoder")
		return []byte{}, false
	}

	return data.Bytes(), true
}

func EncodeBrotli(input []byte) ([]byte, bool) {
	data := bytes.NewBuffer(nil)

	enc := brotli.NewWriterV2(data, brotli.DefaultCompression)

	_, err := enc.Write(input)
	if err != nil {
		fmt.Println("failed to compress brotli")
		return []byte{}, false
	}

	err = enc.Close()
	if err != nil {
		fmt.Println("failed to close brotli encoder")
		return []byte{}, false
	}

	return data.Bytes(), true
}
