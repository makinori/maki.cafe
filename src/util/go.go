package util

import (
	"log/slog"
	"runtime"
	"strings"
)

type GoVersion struct {
	Full       string
	Version    string
	GreenTeaGC bool
}

func GetGoVersion() GoVersion {
	var out GoVersion

	out.Full = runtime.Version()

	if strings.Contains(out.Full, "go1.26") {
		slog.Warn("greenteagc should be default by now! check to make sure")
	}

	out.GreenTeaGC = strings.Contains(out.Full, "greenteagc")

	i := strings.Index(out.Full, " ")
	if i == -1 {
		out.Version = out.Full
	} else {
		out.Version = out.Full[:i]
	}

	return out
}
