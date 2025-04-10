package ui

import (
	_ "embed"

	"github.com/makinori/maki.cafe/common"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed layout.scss
var styles string

func Layout(r *RenderContext, children ...Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=0.85")),
				TitleEl(Text(common.ConfigTitle)),
				Meta(Name("description"), Content(common.ConfigDescription)),
				SCSSEl(r, styles),
			),
			Body(append(
				children,
				JSEl(r),
			)...),
		),
	)
}
