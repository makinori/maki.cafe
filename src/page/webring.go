package page

import (
	"strings"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func webringIcon(filename string, url string) Node {
	style := "image-rendering: pixelated; overflow: hidden;"
	style += "font-size: 14px; display: inline-flex;"
	style += "align-items: center; justify-content: center;"
	style += "background-color: hsl(0deg, 0%, 15%); font-weight: 500;"
	style += "padding: 0;"

	var attrs []Node

	if strings.HasPrefix(filename, "!") {
		// not an image
		attrs = append(attrs, Text(filename[1:]))
		style += "width: 84px; height: 27px; border: solid 2px hsl(0deg, 0%, 25%);"
	} else {
		style += "width: 88px; height: 31px;"
		style += "background-size: 100% auto; background-position: 0 0;"
		style += "background-image: url(\"/webring/" + filename + "\")"
	}

	attrs = append(attrs, Style(style), Href(url))

	return A(attrs...)
}

// 88x31

func Webring() Group {
	return Group{
		H2(Text("friends")),
		Br(),
		Div(
			Style("display: inline-grid; grid-gap: 8px; grid-template-columns: repeat(3, 1fr);"),
			webringIcon("micaela.gif", "https://micae.la"),
			webringIcon("kneesox.png", "https://kneesox.moe"),
			webringIcon("!cmtaz.net", "https://cmtaz.net"),
			webringIcon("!lemon.horse", "https://lemon.horse"),
			webringIcon("!ironsm4sh.nl", "https://ironsm4sh.nl"),
			webringIcon("!0x0a.de", "https://0x0a.de"),
			webringIcon("!pony.best", "https://pony.best"),
		),
		Br(),
		Br(),
		H2(Text("other")),
		Br(),
		Div(
			Style("display: inline-grid; grid-gap: 8px; grid-template-columns: repeat(3, 1fr);"),
			webringIcon("yno.png", "https://ynoproject.net"),
		),
		Br(),
		Br(),
		H2(Text("pony")),
		Br(),
		Div(
			Style("display: inline-grid; grid-gap: 8px; grid-template-columns: repeat(3, 1fr);"),
			webringIcon("anonfilly.png", "https://anonfilly.horse"),
			webringIcon("!pony.town", "https://pony.town"),
			webringIcon("!wetmares", "https://wetmares.org"),
		),
	}
}
