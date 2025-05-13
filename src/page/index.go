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

type link struct {
	Name  string
	Color string
	URL   string
	Icon  string
	Muted bool
	Break bool
}

func makeLinks(links []link) Group {
	var output Group

	for _, item := range links {
		if item.Break {
			output = append(output, Div(Style("height:16px")))
			continue
		}

		style := "background:" + item.Color + ";"
		if item.Muted {
			style += "color:#000;"
		}

		params := Group{
			Style(style),
			Href(item.URL),
		}

		if item.Icon != "" {
			params = append(params, Img(Src(item.Icon)))
		}

		params = append(params, Text(item.Name))

		output = append(output, A(params))
	}

	return output
}

func Index() Group {
	// emailHref, emailTitle, emailJS := botSafeHref("mailto:", common.Email)
	// xmppHref, xmppTitle, xmppJS := botSafeHref("xmpp:", common.XMPP)

	workedOn := makeLinks([]link{
		{
			Name:  "tivoli cloud vr",
			Color: "#e91e63",
			URL:   "https://github.com/tivolicloud",
			Icon:  "/icons/tivoli.svg",
		},
		{
			Name:  "blahaj quest",
			Color: "#3c8ea7",
			URL:   "https://blahaj.quest",
			Icon:  component.EmojiURL("🦈", "noto"),
		},
		{
			Name:  "baltimare leaderboard",
			Color: "#689F38",
			URL:   "https://baltimare.hotmilk.space",
			Icon:  "/icons/happy-anonfilly.png",
		},
		{
			Name:  "melonds metroid hunters",
			Color: "#dd2e44",
			URL:   common.GitHubURL + "/melonPrimeDS",
			Icon:  "/icons/metroid.png",
		},
		{Break: true},
		{
			Name:  "old page",
			Color: "#fff",
			URL:   "https://old.maki.cafe",
			Muted: true,
		},
		{
			Name:  "dots",
			Color: "#fff",
			URL:   common.GitHubURL + "/dots",
			Icon:  "/icons/arch.svg",
			Muted: true,
		},
	})

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
		Br(),
		H1(Text("worked on")),
		Br(),
		Div(
			Style("display: flex; flex-direction: column; align-items: flex-start; gap: 8px"),
			workedOn,
		),
		Br(),
		// Div(
		// 	Style("display: flex; flex-direction: row; gap: 8px"),
		// 	A(
		// 		Class("muted"),
		// 		Text("old page"),
		// 		Href("https://old.maki.cafe"),
		// 	),
		// 	A(
		// 		Class("muted"),
		// 		Href(common.GitHubURL+"/dots"),
		// 		Img(Src("/icons/arch.svg")),
		// 		Text("dots"),
		// 	),
		// ),
	}
}
