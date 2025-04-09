package ui

import (
	. "maragu.dev/gomponents"
)

type Renderer struct {
	PageSCSS map[string]string
}

func Render(page func(*Renderer) Node) string {
	ensureSass()

	r := Renderer{
		PageSCSS: map[string]string{},
	}

	html := Group{
		Layout(&r, page(&r)),
	}.String()

	return html
}
