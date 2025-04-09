package ui

import (
	"github.com/makinori/maki.cafe/shared"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Layout(r *Renderer, children ...Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=0.85")),
				TitleEl(Text(shared.ConfigTitle)),
				Meta(Name("description"), Content(shared.ConfigDescription)),
				SCSSEl(r),
			),
			Body(children...),
		),
	)
}
