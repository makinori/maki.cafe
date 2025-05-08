package page

import (
	"github.com/makinori/maki.cafe/src/common"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index() Group {
	return Group{
		H1(
			Class("title"),
			Span(
				Text("mak"),
				Span(Style("letter-spacing: -4px"), Text("i")),
			),
			Img(Class("pony"), Src("pony.png")),
		),
		Text("software engineer"),
		Br(),
		Text("game developer"),
		Br(),
		Text("server admin"),
		Br(),
		Br(),
		A(
			Href(util.EscapedHTML(common.Email)),
			Title(common.Email),
			Text("email"),
		),
		Text(" "),
		A(
			Href(util.EscapedHTML(common.XMPP)),
			Title(common.XMPP),
			Text("xmpp"),
		),
		Text(" "),
		A(
			Href(common.GitHubURL),
			Title("@"+common.GitHubUsername),
			Text("github"),
		),
		Br(),
		Br(),
		Text("reworking my site..."),
		Br(),
		Br(),
		Text("find my "),
		A(
			Href("https://old.maki.cafe"),
			Text("old page here"),
		),
	}
}
