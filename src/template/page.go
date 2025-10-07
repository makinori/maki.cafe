package template

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"

	"github.com/makinori/goemo"
	"maki.cafe/src/config"
	"maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed styles.scss
	stylesSCSS string
	//go:embed font-maple-mono-maki.scss
	fontSCSS string

	usingIPv6Key string = "usingIPv6"
)

func RenderPage(
	pageFn func(context.Context) Group,
	r *http.Request,
) (string, error) {
	ctx := goemo.InitContext(context.Background())

	ip := util.HTTPGetIPAddress(r)
	if util.IsValidIPv6(ip) {
		ctx = context.WithValue(ctx, usingIPv6Key, true)
	}

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
		Div(Class(goemo.SCSS(ctx, `
			position: absolute;
			margin: auto;
			top: 0;
			left: 0;
			right: 0;
			height: 8px;
			background-color: #ff1744;
		`))),
		pageHeader(ctx, pageHeaderInfo, r),
		pageFn(ctx),
		pageFooter(ctx, currentPagePath),
	)

	pageSCSS := goemo.GetPageSCSS(ctx)

	finalCSS, err := goemo.RenderSCSS(stylesSCSS+"\n"+pageSCSS,
		goemo.SassImport{Filename: "font.scss", Content: fontSCSS},
	)

	if err != nil {
		return "", err
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
		return "", err
	}

	// minify here. dont need to cause scss and gomponents minify

	return siteBuf.String(), nil
}
