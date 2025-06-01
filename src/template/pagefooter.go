package template

import (
	"context"

	"maki.cafe/src/component"
	"maki.cafe/src/config"
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

	return Group{
		Div(Style("margin-top: 100px")),
		hr,
		Div(
			Class("page-footer-pages"),

			P(Text("/"), subPageStyle),
			footerLink(currentPagePath, "/", "index"),
			footerLink(currentPagePath, "/webring", "webring"),
			Div(Class("break")),

			P(Text("/fav/"), subPageStyle),
			footerLink(currentPagePath, "/fav/anime", "anime"),
			footerLink(currentPagePath, "/fav/games", "games"),
		),
		hr,
		Div(
			Class("page-footer-pages"),
			footerLink("", config.GitHubURL+"/maki.cafe", "source code"),
		),
		Br(),
		component.MoeCounter(ctx),
	}
}
