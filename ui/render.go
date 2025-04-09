package ui

import (
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

	html := Group{
		Layout(&r, page(&r)),
	}.String()

	return html
}
