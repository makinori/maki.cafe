package ui

import (
	. "maragu.dev/gomponents"
)

type RenderContext struct {
	SCSS map[string]string
	JS   map[string]string
}

func Render(page func(*RenderContext) Node) string {
	ensureSass()

	r := RenderContext{
		SCSS: map[string]string{},
		JS:   map[string]string{},
	}

	html := Group{
		Layout(&r, page(&r)),
	}.String()

	return html
}
