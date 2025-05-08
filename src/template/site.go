package template

import (
	_ "embed"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed style.css
var siteStyleCSS string

func Site(page Group) Group {
	return Group{Doctype(
		HTML(
			Head(
				Title("maki.cafe"),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=0.6"),
				),
				StyleEl(Raw(siteStyleCSS)),
			),
			Body(
				Div(Class("page-top-strip")),
				page,
				// footer
			),
		),
	)}
}
