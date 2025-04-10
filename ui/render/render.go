package render

import (
	. "maragu.dev/gomponents"
)

type RenderContext struct {
	SCSS map[string]string
	JS   map[string]string // pls use sparingly
}

func Render(
	layout func(*RenderContext, ...Node) Node,
	page func(*RenderContext) Node,
) string {
	ensureSass()

	r := RenderContext{
		SCSS: map[string]string{},
		JS:   map[string]string{},
	}

	html := Group{
		layout(&r, page(&r)),
	}.String()

	return html
}
