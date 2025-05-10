package component

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type PageHeaderInfo struct {
	Big      bool
	PagePath string
}

func PageHeader(info PageHeaderInfo) Group {
	if info.Big {
		return Group{H1(
			Class("page-header-title"),
			Span(
				Text("mak"),
				Span(Style("letter-spacing: -4px"), Text("i")),
			),
			Img(Class("pony"), Src("pony.png")),
		)}
	}

	// small header for all other pages

	return Group{
		Div(
			Class("page-header-small"),
			A(
				Href("/"),
				H1(
					Text("mak"),
					Span(Style("letter-spacing: -4px"), Text("i")),
				),
			),
			A(
				Href(info.PagePath),
				H2(Text(info.PagePath)),
			),
			Div(Style("flex-grow:1")),
			Img(Src("/pony.png")),
		),
		Hr(Class("page-header-small-hr")),
		Br(),
	}

}
