package template

import (
	"context"

	"github.com/makinori/goemo"
	"maki.cafe/src/component"
	"maki.cafe/src/config"
	"maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func footerLink(
	currentPagePath string, pagePath string, name string, prefixNodes ...Node,
) Node {
	props := Group{}

	if currentPagePath == pagePath {
		props = append(props,
			// Text(fmt.Sprintf("{%s}", name)),
			Text(name),
			Class("muted active"),
		)
	} else {
		props = append(props,
			Text(name), Class("muted"), Href(pagePath),
		)
	}

	return A(Group(prefixNodes), props)
}

func pageFooter(ctx context.Context, currentPagePath string) Group {
	// not ready to show footer on index page yet
	// if currentPagePath == "/" {
	// 	return Group{}
	// }

	hr := Hr(Style("width: 250px"))
	subPageStyle := Style("margin-top: 3px;")

	pagesClass := goemo.SCSS(ctx, `
		margin-top: 8px;
		margin-bottom: 8px;
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		column-gap: 8px;
		row-gap: 4px;

		.break {
			flex-basis: 100%;
		}

		// p {
		// 	letter-spacing: -1px;
		// }

		a.active {
			font-weight: 700;
			// border-bottom: solid 2px white;
			// margin-bottom: -2px;
			background: transparent;
			color: #fff;
			// padding-left: 3px;
			// padding-right: 3px;
			// border-radius: 0;
			position: relative;

			&::after {
				content: "";
				margin: auto;
				position: absolute;
				bottom: -2px;
				left: 0;
				right: 0;
				height: 2px;
				background-color: #fff;
				border-radius: 999px;
			}
		}
	`)

	return Group{
		Div(Style("margin-top: 100px")),
		hr,
		Div(
			Class(pagesClass),

			P(Text("/"), subPageStyle),
			footerLink(currentPagePath, "/", "index"),
			footerLink(currentPagePath, "/webring", "webring"),
			Div(Class("break")),

			P(Text("/fav/"), subPageStyle),
			footerLink(currentPagePath, "/fav/anime", "anime"),
			footerLink(currentPagePath, "/fav/games", "games"),
			// Div(Class("break")),

			// P(Text("/past/"), subPageStyle),
			// footerLink(currentPagePath, "/past/avatars", "avatars"),
		),
		hr,
		Div(
			Class(pagesClass),
			Style("align-items: center"),
			footerLink("", config.GitHubURL+"/maki.cafe", "source code"),
			P(
				Style("font-size: 0.8em; margin-left: 4px"),
				Text(util.GetGoVersion()+", {{.RenderTime}}"),
			),
		),
		Br(),
		component.MoeCounter(ctx),
	}
}
