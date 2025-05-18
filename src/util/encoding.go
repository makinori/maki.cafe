package util

import (
	"bytes"

	"github.com/andybalholm/brotli"
	"github.com/klauspost/compress/zstd"
)

func EncodeZstd(input []byte) ([]byte, error) {
	data := bytes.NewBuffer(nil)

	enc, err := zstd.NewWriter(data)
	if err != nil {
		return []byte{}, err
	}

	_, err = enc.Write(input)
	if err != nil {
		return []byte{}, err
	}

	err = enc.Close()
	if err != nil {
		return []byte{}, err
	}

	return data.Bytes(), nil
}

func EncodeBrotli(input []byte) ([]byte, error) {
	data := bytes.NewBuffer(nil)

	enc := brotli.NewWriterV2(data, brotli.DefaultCompression)

	_, err := enc.Write(input)
	if err != nil {
		return []byte{}, err
	}

	err = enc.Close()
	if err != nil {
		return []byte{}, err
	}

	return data.Bytes(), nil
}
