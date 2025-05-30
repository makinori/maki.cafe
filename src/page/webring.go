package page

import (
	"context"
	"strings"

	"maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

// TODO: use grid component

func webringIcon(ctx context.Context, filename string, url string) Node {

	attrs := Group{
		Href("https://" + url),
		Title(url),
	}

	if strings.HasPrefix(filename, "!") {
		// not an image
		attrs = append(attrs, Text(filename[1:]), Classes{
			render.SCSS(ctx, `
				width: 84px;
				height: 27px;
				border: solid 2px hsl(0deg, 0%, 20%);
				align-items: center;
				justify-content: center;
			`): true,
		})
	} else {
		attrs = append(attrs,
			// Style(fmt.Sprintf(
			// 	`background-image: url("/webring/%s")`, filename,
			// )),
			Classes{render.SCSS(ctx, `
				width: 88px;
				height: 31px;
				background-size: 100% auto;
				background-position: 0 0;
				// some are 32 height, move to top
				align-items: flex-start;
				justify-content: flex-start;
				> img {
					width: 88px;
					
					height: auto; 
				}
			`): true},
			Img(Src("/webring/"+filename)),
		)
	}

	return A(attrs)
}

func Webring(ctx context.Context) Group {
	gridClass := render.SCSS(ctx, `
		display: inline-grid;
		grid-gap: 8px;
		grid-template-columns: repeat(4, 1fr);

		> a {
			image-rendering: pixelated;
			overflow: hidden;
			font-size: 14px;
			display: inline-flex;
			background-color: hsl(0deg, 0%, 10%);
			font-weight: 600;
			padding: 0;
		}
	`)

	return Group{
		H2(Text("friends")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "micaela.gif", "micae.la"),
			webringIcon(ctx, "lemon.png", "lemon.horse"),
			webringIcon(ctx, "cmtaz.png", "cmtaz.net"),
			webringIcon(ctx, "!0x0a.de", "0x0a.de"),
			webringIcon(ctx, "skyn3t.gif", "skyn3t.lol"),
			webringIcon(ctx, "!ironsm4sh.nl", "ironsm4sh.nl"),
			Div(),
			webringIcon(ctx, "kneesox.png", "kneesox.moe"),
			// webringIcon(ctx, "!pony.best", "pony.best"),
			webringIcon(ctx, "kayla.gif", "kayla.moe"),
		),
		Br(),
		Br(),
		H2(Text("other")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "yno.png", "ynoproject.net"),
			webringIcon(ctx, "anonfilly.png", "anonfilly.horse"),
		),
		Br(),
		Br(),
		P(Text("feel free to use my button")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "maki.gif", "maki.cafe"),
		),
		// Br(),
		// Br(),
		// H2(Text("pony")),
		// Br(),
		// Div(
		// 	Class(gridClass),
		// 	webringIcon(ctx, "!pony.town", "pony.town"),
		// 	webringIcon(ctx, "!wetmares", "wetmares.org"),
		// ),
	}
}
