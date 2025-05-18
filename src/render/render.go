package render

import (
	"bytes"
	"context"
	_ "embed"

	"github.com/tdewolff/minify/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed style.scss
	styleSCSS string
	//go:embed fonts.scss
	fontsSCSS string

	minifier *minify.M
)

func InitMinifier() {
	// minifier = minify.New()
	// minifier.Add("text/css", &css.Minifier{})
	// minifier.Add("text/html", &html.Minifier{
	// 	KeepDocumentTags:    true,
	// 	KeepQuotes:          true,
	// 	KeepDefaultAttrVals: true,
	// })
}

func Render(
	pageFn func(context.Context) Group,
	currentPagePath string,
) ([]byte, error) {
	ctx := initContext()

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

	body := Body(
		Class(bodyClass),
		Div(Class(SCSS(ctx, `
			position: fixed;
			margin: auto;
			top: 0;
			left: 0;
			right: 0;
			height: 8px;
			background-color: #ff1744;
		`))),
		pageHeader(pageHeaderInfo),
		pageFn(ctx),
		pageFooter(currentPagePath),
	)

	pageSCSS := getPageSCSS(ctx)

	finalCSS, err := renderSass(styleSCSS+"\n"+pageSCSS,
		SassImport{Filename: "fonts.scss", Content: fontsSCSS},
	)

	if err != nil {
		return []byte{}, err
	}

	site := Group{Doctype(
		HTML(
			Head(
				TitleEl(Text(title)),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=0.6"),
				),
				StyleEl(Raw(finalCSS)),
			),
			body,
		),
	)}

	siteBuf := bytes.NewBuffer(nil)
	err = site.Render(siteBuf)
	if err != nil {
		return []byte{}, err
	}

	// minSiteBuf := bytes.NewBuffer(nil)
	// err = minifier.Minify("text/html", minSiteBuf, siteBuf)
	// if err != nil {
	// 	return []byte{}, err
	// }
	// return minSiteBuf.Bytes(), nil

	return siteBuf.Bytes(), nil
}
