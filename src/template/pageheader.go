package template

import (
	"net/http"

	"maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type pageHeaderInfo struct {
	Big      bool
	PagePath string
}

func pageHeader(info pageHeaderInfo, r *http.Request) Group {
	// make sure to set height to avoid flickering
	ponyImg := Img(Class("pony"), Src("/images/pony.png"))

	// dont want to add css to hide incase browsers dont render it

	notabotPixel := Img(Src("/notabot.gif?" + util.NotabotEncode(r)))

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
				notabotPixel,
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
			notabotPixel,
			ponyImg,
		),
		Hr(Style("height: 3px; margin-bottom: 28px")),
	}

}
