package ui

import (
	"strings"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Classes(v []string) Node {
	return Class(strings.Join(v, " "))
}
