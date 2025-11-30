package util

import (
	"runtime"
	"strings"
)

func GetGoVersion() string {
	version := runtime.Version()

	i := strings.Index(runtime.Version(), " ")
	if i > -1 {
		return version[:i]
	}

	return version
}
