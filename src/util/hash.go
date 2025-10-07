package util

import "github.com/cespare/xxhash/v2"

// const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base52 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func HashBytes(data []byte) string {
	hash := xxhash.Sum64(data)
	if hash == 0 {
		return string(base52[0])
	}

	var name string
	for hash > 0 {
		name = string(base52[hash%52]) + name
		hash /= 52
	}

	return name
}

func HashString(data string) string {
	return HashBytes([]byte(data))
}
