package template

import (
	"context"
	"net/http"

	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohttp"
	"maki.cafe/src/component"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type pageHeaderInfo struct {
	Big      bool
	PagePath string
}

func pageHeader(ctx context.Context, info pageHeaderInfo, r *http.Request) Group {
	// make sure to set height to avoid flickering
	ponyImg := Img(Class("maki-pony"), Src("/images/pony-header.png"))

	// dont want to add css to hide incase browsers dont render it

	notabotPixel := Img(Src("/notabot.gif?" + emohttp.NotABotURLQuery(r)))

	if info.Big {
		return Group{
			Div(
				Class(goemo.SCSS(ctx, `
					font-size: 100px;
					font-weight: bold;
					line-height: 100px;
					color: white;
					margin-top: -8px;
					// margin-bottom: 8px;
					gap: 24px;
					display: flex;
					flex-direction: row;
					align-items: flex-end;

					> svg {
						margin-bottom: 12px;
					}

					> .maki-pony {
						height: 128px;
						margin-top: -24px;
					}

					@media (max-width: $page-break-width) {
						margin-top: 0px;
						padding-top: 24px;
					}
				`)),
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
			Class(goemo.SCSS(ctx, `
				display: flex;
				flex-direction: row;
				align-items: flex-end;
				// gap: 8px;
				margin-top: -16px;

				@media (max-width: $page-break-width) {
					margin-top: -8px;
				}

				svg {
					margin-right: 6px;
					margin-bottom: 8px;
				}

				h2 {
					font-size: 28px;
					margin-bottom: 3px;
					font-weight: 600;
					letter-spacing: 0px;
					// opacity: 0.5;
				}

				> a {
					background-color: transparent;
					color: #fff;
					padding: inherit;
				}

				> .maki-pony {
					// height: 64px;
					height: 80px;
				}
			`)),
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
