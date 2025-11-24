package util

import (
	"runtime"
)

func GetGoVersion() string {
	version := runtime.Version()
	return version
	// i := strings.Index(version, " ")
	// if i == -1 {
	// 	return version
	// } else {
	// 	return version[:i]
	// }
}
