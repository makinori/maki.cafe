package page

import (
	"github.com/makinori/maki.cafe/src/common"
	"github.com/makinori/maki.cafe/src/component"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// func botSafeHref(prefix string, value string) (string, string, string) {
// 	// this upsets me

// 	// encoded := base64.StdEncoding.EncodeToString([]byte(value))
// 	encoded := strings.ReplaceAll(value, "@", "[at]")
// 	encoded = strings.ReplaceAll(encoded, ".", "[dot]")

// 	// keep it legible. even though this only happens when js is disabled,
// 	// the whole point is to make sure the user can get the address
// 	jsHref := fmt.Sprintf(`javascript:alert("%s")`, encoded)

// 	js := strings.Join([]string{
// 		`{`,
// 		`let e = document.currentScript.parentElement;`,
// 		`e.title = e.title.replaceAll("[at]","@").replaceAll("[dot]",".");`,
// 		`e.href = "` + prefix + `"+e.title;`,
// 		`}`,
// 	}, " ")

// 	return jsHref, encoded, js
// }

type workedOn struct {
	Name  string
	Color string
	URL   string
	Icon  string
}

func Index() Group {
	// emailHref, emailTitle, emailJS := botSafeHref("mailto:", common.Email)
	// xmppHref, xmppTitle, xmppJS := botSafeHref("xmpp:", common.XMPP)

	workedOn := []workedOn{
		workedOn{
			Name:  "tivoli cloud vr",
			Color: "#e91e63",
			URL:   "https://github.com/tivolicloud",
			Icon:  "/icons/tivoli.svg",
		},
		workedOn{
			Name:  "blahaj quest",
			Color: "#3c8ea7",
			URL:   "https://blahaj.quest",
			Icon:  component.EmojiURL("🦈", "noto"),
		},
		workedOn{
			Name:  "baltimare leaderboard",
			Color: "#689F38",
			URL:   "https://baltimare.hotmilk.space",
			Icon:  "/icons/happy-anonfilly.png",
		},
		workedOn{
			Name:  "melonds metroid hunters",
			Color: "#dd2e44",
			URL:   common.GitHubURL + "/melonPrimeDS",
			Icon:  "/icons/metroid.png",
		},
	}

	var workedOnNodes Group

	for _, item := range workedOn {
		workedOnNodes = append(workedOnNodes, A(
			Style("background:"+item.Color),
			Href(item.URL),
			Img(Src(item.Icon)),
			Text(item.Name),
		))
	}

	return Group{
		H3(Text("software engineer")),
		H3(Text("game developer")),
		H3(Text("server admin")),
		Br(),
		A(
			Text("email"),
			Title(common.Email),
			Href("/email"),
			// &util.AttrRaw{
			// 	Name:  "href",
			// 	Value: "mailto:" + util.EscapedHTML(common.Email),
			// },
			// Href(emailHref),
			// Script(Raw(emailJS)),
		),
		Text(" "),
		A(
			Text("xmpp"),
			Title(common.XMPP),
			Href("/xmpp"),
			// &util.AttrRaw{
			// 	Name:  "href",
			// 	Value: "xmpp:" + util.EscapedHTML(common.XMPP),
			// },
			// Href(xmppHref),
			// Script(Raw(xmppJS)),
		),
		Text(" "),
		A(
			Href(common.GitHubURL),
			Title("@"+common.GitHubUsername),
			Text("github"),
		),
		Br(),
		// Br(),
		// Text("find my "),
		// A(
		// 	Class("muted"),
		// 	Href("https://old.maki.cafe"),
		// 	Text("old page here"),
		// ),
		// Br(),
		Br(),
		Br(),
		H1(Text("worked on")),
		Br(),
		Div(
			Style("display: flex; flex-direction: column; align-items: flex-start; gap: 8px"),
			workedOnNodes,
		),
	}
}
