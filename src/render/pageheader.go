package render

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type pageHeaderInfo struct {
	Big      bool
	PagePath string
}

func pageHeader(info pageHeaderInfo) Group {
	// make sure to set height to avoid flickering
	ponyImg := Img(Class("pony"), Src("/images/pony.png"))

	if info.Big {
		return Group{
			H1(
				Class("page-header-title"),
				Span(
					Text("ma"),
					Span(Class("k"), Text("k")),
					Span(Class("i"), Text("i")),
				),
				ponyImg,
			),
			Hr(
				Class("page-header-hr"),
				Style("height: 4px; width: 360px; margin-bottom: 18px"),
			),
			// Br(),
		}
	}

	// small header for all other pages

	return Group{
		Div(
			Class("page-header-small"),
			A(
				Href("/"),
				H1(
					Text("ma"),
					Span(Class("k"), Text("k")),
					Span(Class("i"), Text("i")),
				),
			),
			A(
				Href(info.PagePath),
				H2(Text(info.PagePath)),
			),
			Div(Style("flex-grow:1")),
			ponyImg,
		),
		Hr(Style("height: 3px; margin-bottom: 28px")),
	}

}
