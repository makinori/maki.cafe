package page

import (
	"context"
	"strings"

	"github.com/makinori/foxlib/foxcss"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TODO: use grid component
// TODO: optimize use of render sscss

func webringIcon(
	ctx context.Context, filename string, url string, webringAttrs ...string,
) Node {
	attrs := Group{}

	if url != "" {
		attrs = append(attrs, Href("https://"+url))
	}

	classesSuffix := ""
	hasTitle := false

	for _, webringAttr := range webringAttrs {
		if webringAttr == "transparent" {
			classesSuffix += " transparent"
		} else if strings.HasPrefix(webringAttr, "title:") {
			attrs = append(attrs, Title(webringAttr[6:]))
			hasTitle = true
		}
	}

	if !hasTitle {
		attrs = append(attrs, Title(url))
	}

	if strings.HasPrefix(filename, "!") {
		// not an image
		attrs = append(attrs, Text(filename[1:]), Class(
			foxcss.Class(ctx, `
				width: 84px;
				height: 27px;
				border: solid 2px hsl(0deg, 0%, 20%);
				align-items: center;
				justify-content: center;
				background-color: hsl(0deg, 0%, 10%);
			`)+classesSuffix),
		)
	} else {
		attrs = append(attrs,
			// Style(fmt.Sprintf(
			// 	`background-image: url("/webring/%s")`, filename,
			// )),
			Class(foxcss.Class(ctx, `
				width: 88px;
				height: 31px;
				background-color: #fff;
				background-size: 100% auto;
				background-position: 0 0;
				// some are 32 height, move to top
				align-items: flex-start;
				justify-content: flex-start;
				> img {
					margin: 0;
					width: 88px;
					height: auto; 
				}
				&.transparent {
					background: transparent;
				}
			`)+classesSuffix),
			Img(Src("/webring/"+filename)),
		)
	}

	return A(attrs)
}

func sillyWebring(ctx context.Context, gridClass string) Group {
	return Group{
		H2(Text("silly")),
		Br(),
		Div(
			Class(gridClass+" wider"),
			webringIcon(ctx, "silly/agplv3.png", "fsf.org"),
			webringIcon(ctx, "silly/amd.gif", "amd.com"),
			webringIcon(ctx, "silly/blender.gif", "blender.org"),
			webringIcon(ctx, "silly/docker.png", "docker.com"),
			webringIcon(ctx, "silly/forgejo.png", "forgejo.org"),
			webringIcon(ctx, "silly/gbanet.gif", ""),
			webringIcon(ctx, "silly/gimp.gif", "gimp.org"),
			webringIcon(ctx, "silly/grapheneos.gif", "grapheneos.org"),
			webringIcon(ctx, "silly/halflife.gif", "store.steampowered.com"),
			webringIcon(ctx, "silly/imhex.png", "imhex.werwolv.net"),
			webringIcon(ctx, "silly/imissxp.gif", ""),
			webringIcon(ctx, "silly/implementrssnow.png", "miniflux.app"),
			webringIcon(ctx, "silly/inkscape.svg", "", "transparent"),
			webringIcon(ctx, "silly/internetarchive.gif", "archive.org"),
			webringIcon(ctx, "silly/jellyfin.gif", "jellyfin.org"),
			webringIcon(ctx, "silly/kagi.png", "kagi.com"),
			webringIcon(ctx, "silly/learnhtml.gif", ""),
			webringIcon(ctx, "silly/linuxnow2.gif", ""),
			webringIcon(ctx, "silly/nvidia.gif", "nvidia.com"),
			webringIcon(ctx, "silly/opengl.gif", "learnopengl.com"),
			webringIcon(ctx, "silly/qbittorrent.png", "qbittorrent.org"),
			webringIcon(ctx, "silly/sdl.gif", "libsdl.org"),
			webringIcon(ctx, "silly/secondlife.gif", "secondlife.com"),
			webringIcon(ctx, "silly/solitaire.png", "store.steampowered.com/app/1988540"),
			webringIcon(ctx, "silly/steam.gif", "store.steampowered.com"),
			webringIcon(ctx, "silly/tor.gif", "torproject.org"),
			webringIcon(ctx, "silly/traderjoes.gif", "traderjoes.com"),
			webringIcon(ctx, "silly/ubo.png", "ublockorigin.com"),
			webringIcon(ctx, "silly/wii.gif", "nintendo.com"),
			webringIcon(ctx, "silly/xmpp.gif", "xmpp.org"),
			webringIcon(ctx, "silly/yosemite.gif", "nps.gov/yose"),
			webringIcon(ctx, "silly/yumenikki.gif", "ynoproject.net"),
		),
		Br(),
		Br(),
	}
}

func Webring(ctx context.Context) Group {
	gridClass := foxcss.Class(ctx, `
		display: inline-grid;
		grid-gap: 8px;
		grid-template-columns: repeat(4, 1fr);

		&.wider {
			grid-template-columns: repeat(5, 1fr);
		}

		> a {
			image-rendering: pixelated;
			overflow: hidden;
			font-size: 11px;
			display: inline-flex;
			font-weight: 600;
			padding: 0;
			border-radius: 0;
		}
	`)

	return Group{
		H2(Text("friends")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "0x0ade.gif", "0x0a.de"),
			webringIcon(ctx, "lemon.png", "lemon.horse"),
			webringIcon(ctx, "cmtaz.png", "cmtaz.net"),
			webringIcon(ctx, "micaela.gif", "micae.la"),
			webringIcon(ctx, "kayla.gif", "kayla.moe"),
			webringIcon(ctx, "kneesox.png", "kneesox.moe"),
			webringIcon(ctx, "!ironsm4sh.nl", "ironsm4sh.nl"),
			webringIcon(ctx, "skyn3t.gif", "skyn3t.lol"),
			// webringIcon(ctx, "!pony.best", "pony.best"),
			// Div(),
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
		// https://cyber.dabamos.de/88x31/
		// https://capstasher.neocities.org/88x31collection
		// TODO: https://neonaut.neocities.org/cyber/88x31
		Br(),
		Br(),
		// sillyWebring(ctx, gridClass),
		P(Text("feel free to use my button")),
		P(Text("although it needs to be updated")),
		Br(),
		Div(
			Class(gridClass),
			webringIcon(ctx, "maki.gif", "maki.cafe", "title:or use maki@2x.gif"),
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
