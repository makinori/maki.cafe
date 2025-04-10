package render

import (
	. "maragu.dev/gomponents"
)

type RenderContext struct {
	SCSS   map[string]string
	HeadJS map[string]string
	BodyJS map[string]string
}

func Render(
	layout func(*RenderContext, ...Node) Node,
	page func(*RenderContext) Node,
) string {
	ensureSass()

	r := RenderContext{
		SCSS:   map[string]string{},
		HeadJS: map[string]string{},
		BodyJS: map[string]string{},
	}

	html := Group{
		layout(&r, page(&r)),
	}.String()

	return html
}
