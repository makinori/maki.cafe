package template

import (
	_ "embed"

	"github.com/makinori/maki.cafe/src/component"
	"github.com/makinori/maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed style.scss
	styleSCSS string
	//go:embed fonts.scss
	fontsSCSS string
)

func Site(page Group, currentPagePath string) (Group, error) {
	finalSCSS, err := render.RenderSass(styleSCSS,
		render.SassImport{Filename: "fonts.scss", Content: fontsSCSS},
	)

	if err != nil {
		return Group{}, err
	}

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
				StyleEl(Raw(finalSCSS)),
			),
			Body(
				Div(Class("page-top-strip")),
				component.PageHeader(pageHeaderInfo),
				page,
				component.PageFooter(currentPagePath),
			),
		),
	)}, nil
}
