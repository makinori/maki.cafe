package template

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohttp"
	"maki.cafe/src/config"
	"maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	//go:embed utils.scss
	utilsSCSS string
	//go:embed colors.scss
	colorsSCSS string
	//go:embed styles.scss
	stylesSCSS string
	//go:embed font-maple-mono-maki.scss
	fontSCSS string

	// https://gist.github.com/borlaym/585e2e09dd6abd9b0d0a
	// instead compiled my own with https://emojipedia.org/nature#list
	//go:embed animals.json
	animalsJSON []byte
	animals     []string

	usingIPv6Key string = "usingIPv6"
)

func init() {
	err := json.Unmarshal(animalsJSON, &animals)
	if err != nil {
		slog.Error("failed to parse animals json", "err", animals)
		os.Exit(1)
	}
}

func metaTagWithName(name, content string) Node {
	return Meta(Name(name), Content(content))
}

func metaTagWithProperty(name, content string) Node {
	return Meta(Attr("property", name), Content(content))
}

func RenderPage(
	pageFn func(context.Context) Group,
	r *http.Request,
) (string, error) {
	ctx := goemo.InitContext(context.Background())

	ctx = goemo.UseWords(ctx, animals, time.Now().Format(time.DateOnly))

	ip := emohttp.GetIPAddress(r)
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
			// background-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCA5MDAgMTAwIj48cGF0aCBmaWxsPSIjZTY2MTAwIiBkPSJNMCAwaDkwMHYxMDBIMHoiLz48cGF0aCBmaWxsPSIjYzY0NjAwIiBkPSJNMTAwIDBoOTAwdjEwMEgxMDB6Ii8+PHBhdGggZmlsbD0iI2MxMzgxZCIgZD0iTTIwMCAwaDkwMHYxMDBIMjAweiIvPjxwYXRoIGZpbGw9IiNiYzMzMjMiIGQ9Ik0zMDAgMGg5MDB2MTAwSDMwMHoiLz48cGF0aCBmaWxsPSIjYjgyZTJhIiBkPSJNMzI1IDBoOTAwdjEwMEgzMjV6Ii8+PHBhdGggZmlsbD0iI2IzMjkzMSIgZD0iTTM1MCAwaDkwMHYxMDBIMzUweiIvPjxwYXRoIGZpbGw9IiNhZjI0MzgiIGQ9Ik0zNzUgMGg5MDB2MTAwSDM3NXoiLz48cGF0aCBmaWxsPSIjOTIxZjQ4IiBkPSJNNDAwIDBoOTAwdjEwMEg0MDB6Ii8+PHBhdGggZmlsbD0iIzcwMjM0ZSIgZD0iTTUwMCAwaDkwMHYxMDBINTAweiIvPjxwYXRoIGZpbGw9IiM2NzIzNGQiIGQ9Ik02MDAgMGg5MDB2MTAwSDYwMHoiLz48cGF0aCBmaWxsPSIjNWYyNDRjIiBkPSJNNjI1IDBoOTAwdjEwMEg2MjV6Ii8+PHBhdGggZmlsbD0iIzU2MjQ0YiIgZD0iTTY1MCAwaDkwMHYxMDBINjUweiIvPjxwYXRoIGZpbGw9IiM0ZTI1NGEiIGQ9Ik02NzUgMGg5MDB2MTAwSDY3NXoiLz48cGF0aCBmaWxsPSIjMzAyMjNiIiBkPSJNNzAwIDBoOTAwdjEwMEg3MDB6Ii8+PHBhdGggZmlsbD0iIzI0MWYzMSIgZD0iTTgwMCAwaDkwMHYxMDBIODAweiIvPjwvc3ZnPg==);
			// background-size: 100%;
		`))),
		pageHeader(ctx, pageHeaderInfo, r),
		pageFn(ctx),
		pageFooter(ctx, currentPagePath),
	)

	pageSCSS := goemo.GetPageSCSS(ctx)

	finalCSS, err := goemo.RenderSCSS(stylesSCSS+"\n"+pageSCSS,
		goemo.SassImport{Filename: "utils.scss", Content: utilsSCSS},
		goemo.SassImport{Filename: "colors.scss", Content: colorsSCSS},
		goemo.SassImport{Filename: "font.scss", Content: fontSCSS},
	)

	if err != nil {
		return "", err
	}

	head := Group{
		TitleEl(Text(title)),
		// primary
		metaTagWithName("title", title),
		metaTagWithName("description", config.Description),
		// open graph
		metaTagWithProperty("og:type", "website"),
		metaTagWithProperty("og:url", "https://maki.cafe"),
		metaTagWithProperty("og:title", title),
		metaTagWithProperty("og:description", config.Description),
		metaTagWithProperty("og:image", config.SiteImage),
		// twitter
		metaTagWithProperty("twitter:card", "summary"), // summary_large_image
		metaTagWithProperty("twitter:url", "https://maki.cafe"),
		metaTagWithProperty("twitter:title", title),
		metaTagWithProperty("twitter:description", config.Description),
		metaTagWithProperty("twitter:image", config.SiteImage),
		// rest
		metaTagWithName("viewport", "width=device-width, initial-scale=0.6"),
		extraHeadNodes,
		StyleEl(Raw(finalCSS)),
	}

	site := Group{Doctype(
		HTML(
			Head(head),
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
