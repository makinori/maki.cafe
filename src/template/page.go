package template

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"

	"maki.cafe/src/config"
	"maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed style.scss
	styleSCSS string
	//go:embed fonts.scss
	fontsSCSS string
)

func RenderPage(
	pageFn func(context.Context) Group,
	r *http.Request,
) ([]byte, error) {
	ctx := render.InitContext()

	currentPagePath := r.URL.Path

	pageHeaderInfo := pageHeaderInfo{
		PagePath: currentPagePath,
	}

	title := "maki.cafe"
	bodyClass := ""
	var extraHeadNodes Group

	if currentPagePath == "/" {
		pageHeaderInfo.Big = true
		bodyClass = "noblur"
		extraHeadNodes = append(extraHeadNodes, Meta(
			Name("go-import"),
			Content("maki.cafe git "+config.GitHubURL+"/maki.cafe"),
		))
	} else {
		title += currentPagePath
	}

	body := Body(
		Class(bodyClass),
		Div(Class(render.SCSS(ctx, `
			position: absolute;
			margin: auto;
			top: 0;
			left: 0;
			right: 0;
			height: 8px;
			background-color: #ff1744;
		`))),
		pageHeader(pageHeaderInfo, r),
		pageFn(ctx),
		pageFooter(ctx, currentPagePath),
	)

	pageSCSS := render.GetPageSCSS(ctx)

	finalCSS, err := render.RenderSass(styleSCSS+"\n"+pageSCSS,
		render.SassImport{Filename: "fonts.scss", Content: fontsSCSS},
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
				extraHeadNodes,
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

	// minify here. dont need to cause scss and gomponents minify

	return siteBuf.Bytes(), nil
}
