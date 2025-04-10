package common

import (
	"encoding/binary"
	"hash/crc32"
	"runtime"
)

func Prepend[T any](n T, rest []T) []T {
	return append([]T{n}, rest...)
}

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HashBytes(data []byte) string {
	hash := crc32.ChecksumIEEE(data)
	if hash == 0 {
		return string(base62[0])
	}

	var name string
	for hash > 0 {
		name = string(base62[hash%62]) + name
		hash /= 62
	}

	return name
}

func HashString(data string) string {
	return HashBytes([]byte(data))
}

func UniqueHashPC() string {
	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get unique hash")
	}

	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(pc))

	return HashBytes(data)
}
