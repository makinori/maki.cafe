package component

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func footerLink(currentPagePath string, pagePath string, name string) Node {
	class := "muted"
	href := pagePath

	if currentPagePath == pagePath {
		class += " active"
		href = "#"
	}

	return A(
		Text(name),
		Class(class),
		Href(href),
	)
}

func PageFooter(currentPagePath string) Group {
	return []Node{
		// need a good hr
		Br(),
		Br(),
		Br(),
		Br(),
		Br(),
		Hr(),
		Div(
			Class("footer"),
			footerLink(currentPagePath, "/", "index"),
			footerLink(currentPagePath, "/awesome", "awesome"),
			footerLink(currentPagePath, "/webring", "webring"),
			Div(Class("break")),
			footerLink("", "https://old.maki.cafe", "old page"),
		),
	}
}
