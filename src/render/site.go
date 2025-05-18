package render

import (
	"context"
	_ "embed"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed style.scss
	styleSCSS string
	//go:embed fonts.scss
	fontsSCSS string
)

func Site(ctx context.Context, page Group, currentPagePath string) (Group, error) {
	pageSCSS := getPageSCSS(ctx)

	finalCSS, err := renderSass(styleSCSS+"\n"+pageSCSS,
		SassImport{Filename: "fonts.scss", Content: fontsSCSS},
	)

	if err != nil {
		return Group{}, err
	}

	pageHeaderInfo := pageHeaderInfo{
		PagePath: currentPagePath,
	}

	title := "maki.cafe"
	bodyClass := ""

	if currentPagePath == "/" {
		pageHeaderInfo.Big = true
		bodyClass = "noblur"
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
				StyleEl(Raw(finalCSS)),
			),
			Body(
				Class(bodyClass),
				Div(Class("page-top-strip")),
				pageHeader(pageHeaderInfo),
				page,
				pageFooter(currentPagePath),
			),
		),
	)}, nil
}
