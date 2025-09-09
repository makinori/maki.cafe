package template

import (
	"net/http"

	"maki.cafe/src/component"
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
			Div(
				Class("page-header-title"),
				component.Maki(
					Attr("fill", "#fff"),
					Height("80"),
				),
				ponyImg,
				notabotPixel,
			),
			Hr(
				Style("height: 4px; width: 335px; margin-bottom: 18px"),
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
				component.Maki(
					Attr("fill", "#fff"),
					Height("50"),
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
