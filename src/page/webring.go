package page

import (
	"context"
	"fmt"
	"strings"

	"github.com/makinori/maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func webringIcon(ctx context.Context, filename string, url string) Node {
	baseClass := render.SCSS(ctx, `
		image-rendering: pixelated;
		overflow: hidden;
		font-size: 14px;
		display: inline-flex;
		align-items: center;
		justify-content: center;
		background-color: hsl(0deg, 0%, 10%);
		font-weight: 600;
		padding: 0;
	`)

	attrs := Group{Href(url)}

	if strings.HasPrefix(filename, "!") {
		// not an image
		attrs = append(attrs, Text(filename[1:]), Classes{
			baseClass: true, render.SCSS(ctx, `
				width: 84px;
				height: 27px;
				border: solid 2px hsl(0deg, 0%, 20%);
			`): true,
		})
	} else {
		attrs = append(attrs, Style(
			fmt.Sprintf(`background-image: url("/webring/%s")`, filename),
		), Classes{
			baseClass: true, render.SCSS(ctx, `
				width: 88px;
				height: 31px;
				background-size: 100% auto;
				background-position: 0 0;
			`): true,
		})
	}

	return A(attrs)
}

func Webring(ctx context.Context) Group {
	gridClass := render.SCSS(ctx, `
		display: inline-grid;
		grid-gap: 8px;
		grid-template-columns: repeat(4, 1fr);
	`)

	return Group{
		H2(Text("friends")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "micaela.gif", "https://micae.la"),
			webringIcon(ctx, "skyn3t.gif", "https://skyn3t.lol"),
			webringIcon(ctx, "!cmtaz.net", "https://cmtaz.net"),
			webringIcon(ctx, "!lemon.horse", "https://lemon.horse"),
			webringIcon(ctx, "kneesox.png", "https://kneesox.moe"),
			webringIcon(ctx, "!ironsm4sh.nl", "https://ironsm4sh.nl"),
			webringIcon(ctx, "!0x0a.de", "https://0x0a.de"),
			webringIcon(ctx, "!pony.best", "https://pony.best"),
			webringIcon(ctx, "kayla.gif", "https://kayla.moe"),
		),
		Br(),
		Br(),
		H2(Text("other")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "yno.png", "https://ynoproject.net"),
			webringIcon(ctx, "anonfilly.png", "https://anonfilly.horse"),
		),
		Br(),
		Br(),
		P(Text("feel free to use my button")),
		Br(),
		webringIcon(ctx, "maki.gif", "https://maki.cafe"),
		// Br(),
		// Br(),
		// H2(Text("pony")),
		// Br(),
		// Div(
		// 	Class(gridClass),
		// 	webringIcon(ctx, "!pony.town", "https://pony.town"),
		// 	webringIcon(ctx, "!wetmares", "https://wetmares.org"),
		// ),
	}
}
