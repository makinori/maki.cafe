package page

import (
	"github.com/makinori/maki.cafe/src/common"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index() Group {
	return Group{
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
		// Br(),
		// Br(),
		// Br(),
		// H2(Text("stuff ive made")),
		// Br(),
		// A(
		// 	Href("https://old.maki.cafe"),
		// 	Style("background: #e91e63"),
		// 	Img(
		// 		Src("/icons/tivoli.svg"),
		// 	),
		// 	Text("tivoli cloud vr"),
		// ),
		// Br(),
		// A(
		// 	Href("https://blahaj.quest"),
		// 	Style("background: #3c8ea7"),
		// 	Img(
		// 		// Src("/icons/blahaj.png"),
		// 		Src(component.EmojiURL("🔍", "noto")),
		// 	),
		// 	Text("blahaj quest"),
		// ),
	}
}
