package page

import (
	"fmt"
	"strings"

	"github.com/makinori/maki.cafe/src/common"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func botSafeHref(prefix string, value string) (string, string, string) {
	// this upsets me

	// encoded := base64.StdEncoding.EncodeToString([]byte(value))
	encoded := strings.ReplaceAll(value, "@", "[at]")
	encoded = strings.ReplaceAll(encoded, ".", "[dot]")

	// keep it legible. even though this only happens when js is disabled,
	// the whole point is to make sure the user can get the address
	jsHref := fmt.Sprintf(`javascript:alert("%s")`, encoded)

	js := strings.Join([]string{
		`{`,
		`let e = document.currentScript.parentElement;`,
		`e.title = e.title.replaceAll("[at]","@").replaceAll("[dot]",".");`,
		`e.href = "` + prefix + `"+e.title;`,
		`}`,
	}, " ")

	return jsHref, encoded, js
}

func Index() Group {
	emailHref, emailTitle, emailJS := botSafeHref("mailto:", common.Email)
	xmppHref, xmppTitle, xmppJS := botSafeHref("xmpp:", common.XMPP)

	return Group{
		H3(Text("software engineer")),
		H3(Text("game developer")),
		H3(Text("server admin")),
		Br(),
		A(
			Href(emailHref),
			Title(emailTitle),
			Text("email"),
			Script(Raw(emailJS)),
		),
		Text(" "),
		A(
			Href(xmppHref),
			Title(xmppTitle),
			Text("xmpp"),
			Script(Raw(xmppJS)),
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
