package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(r *Renderer, children ...Node) Node {
	return HTML(
		Head(
			TitleEl(Text("maki")),
			SCSSEl(r),
		),
		Body(children...),
	)
}
