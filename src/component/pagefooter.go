package component

import (
	"github.com/makinori/maki.cafe/src/common"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func footerLink(currentPagePath string, pagePath string, name string) Node {
	props := []Node{
		Text(name),
	}

	if currentPagePath != pagePath {
		props = append(props, Class("muted"))
		props = append(props, Href(pagePath))
	} else {
		props = append(props, Class("muted active"))
	}

	return A(props...)
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
		Hr(Style("width: 200px")),
		Div(
			Class("page-footer-pages"),
			footerLink(currentPagePath, "/", "index"),
			// footerLink(currentPagePath, "#", "works"),
			footerLink(currentPagePath, "/anime", "/anime"),
			// footerLink(currentPagePath, "#", "games"),
			// footerLink(currentPagePath, "#", "webring"),
			Div(Class("break")),
			footerLink("", common.GitHubURL+"/maki.cafe", "source code"),
			// Div(Class("break")),
			// footerLink("", "https://old.maki.cafe", "old page"),
		),
	}
}
