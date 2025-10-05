package util

import (
	"runtime"
	"strings"
)

func GetGoVersion() string {
	version := runtime.Version()
	i := strings.Index(version, " ")
	if i == -1 {
		return version
	} else {
		return version[:i]
	}
}
