package template

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"

	"github.com/makinori/emgotion"
	"maki.cafe/src/config"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed styles.scss
	stylesSCSS string
	//go:embed font-maple-mono-maki.scss
	fontSCSS string
)

func RenderPage(
	pageFn func(context.Context) Group,
	r *http.Request,
) ([]byte, error) {
	ctx := emgotion.InitContext(context.Background())

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
		Div(Class(emgotion.SCSS(ctx, `
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

	pageSCSS := emgotion.GetPageSCSS(ctx)

	finalCSS, err := emgotion.RenderSCSS(stylesSCSS+"\n"+pageSCSS,
		emgotion.SassImport{Filename: "font.scss", Content: fontSCSS},
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
