package shared

import (
	"hash/crc32"
)

func Prepend[T any](n T, rest []T) []T {
	return append([]T{n}, rest...)
}

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HashString(data string) string {
	hash := crc32.ChecksumIEEE([]byte(data))
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
