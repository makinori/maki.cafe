package template

import (
	_ "embed"

	"github.com/makinori/maki.cafe/src/component"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed style.css
var siteStyleCSS string

func Site(page Group, currentPagePath string) Group {
	pageHeaderInfo := component.PageHeaderInfo{
		PagePath: currentPagePath,
	}

	title := "maki.cafe"

	if currentPagePath == "/" {
		pageHeaderInfo.Big = true
	} else {
		title += currentPagePath
	}

	return Group{Doctype(
		HTML(
			Head(
				TitleEl(Text(title)),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=0.6"),
				),
				StyleEl(Raw(siteStyleCSS)),
			),
			Body(
				Div(Class("page-top-strip")),
				component.PageHeader(pageHeaderInfo),
				page,
				component.PageFooter(currentPagePath),
			),
		),
	)}
}
