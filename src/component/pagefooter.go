package component

import (
	"github.com/makinori/maki.cafe/src/config"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func footerLink(
	currentPagePath string, pagePath string, name string, prefixNodes ...Node,
) Node {
	props := Group{
		Text(name),
	}

	if currentPagePath != pagePath {
		props = append(props, Class("muted"))
		props = append(props, Href(pagePath))
	} else {
		props = append(props, Class("muted active"))
	}

	return A(Group(prefixNodes), props)
}

func PageFooter(currentPagePath string) Group {
	// not ready to show footer on index page yet
	// if currentPagePath == "/" {
	// 	return Group{}
	// }

	var spacing Group
	for range 8 {
		spacing = append(spacing, Br())
	}

	return Group{
		spacing,
		Hr(Style("width: 300px")),
		Div(
			Class("page-footer-pages"),

			footerLink(currentPagePath, "/", "index"),
			// footerLink(currentPagePath, "#", "works"),
			// footerLink(currentPagePath, "#", "webring"),

			// Div(Class("break")),

			// P(Text("/interests")),
			footerLink(currentPagePath, "/anime", "/anime"),
			footerLink(currentPagePath, "/webring", "/webring"),
			// footerLink(currentPagePath, "#", "games"),
		),
		Hr(Style("width: 300px")),
		Div(
			Class("page-footer-pages"),
			footerLink("", config.GitHubURL+"/maki.cafe", "source code"),
			// Div(Class("break")),
			// footerLink("", "https://old.maki.cafe", "old page"),
			// Div(Class("break")),
			// footerLink("", config.GitHubURL+"/dots", "dots", Img(Src("/icons/arch.svg"))),
		),
	}
}
