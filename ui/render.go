package ui

import (
	"runtime"

	. "maragu.dev/gomponents"
)

type Renderer struct {
	SharedSCSS map[string]string
}

func Render(page func(*Renderer) Node) string {
	ensureSass()

	r := Renderer{
		SharedSCSS: map[string]string{},
	}

	runtime.StartTrace()

	html := Group{
		Layout(&r, page(&r)),
	}.String()

	return html
}
