package page

import (
	"context"

	"github.com/makinori/goemo/emohtml"
	"maki.cafe/src/config"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type link struct {
	Name  string
	Color string
	URL   string
	Icon  string
	Muted bool
	Break bool
	Title string
}

func makeLinks(links []link) Group {
	var output Group

	for _, item := range links {
		if item.Break {
			output = append(output, Div(Style("height:16px")))
			continue
		}

		params := Group{
			Href(item.URL),
		}

		if item.Title != "" {
			params = append(params, Title(item.Title))
		}

		style := ""
		if item.Color != "" {
			style += "background:" + item.Color + ";"
		}
		if item.Muted {
			style += "color:#000;"
		}
		if style != "" {
			params = append(params, Style(style))
		}

		if item.Icon != "" {
			params = append(params, Img(Src(item.Icon)))
		}

		params = append(params, Text(item.Name))

		output = append(output, A(params))
	}

	return output
}

func Index(ctx context.Context) Group {
	social := emohtml.VStack(ctx,
		emohtml.HStack(ctx, makeLinks([]link{
			{
				Name:  "email",
				URL:   "/email",
				Title: config.Email,
				// Color: "#333",
			},
			{
				Name:  "xmpp",
				URL:   "/xmpp",
				Title: config.XMPP,
				Color: "#227ee1",
			},
			{
				Name:  "tox",
				URL:   config.ToxURI,
				Title: config.Tox,
				Color: "#ff8f00",
			},
			{
				Name:  "matrix",
				URL:   config.MatrixURL,
				Title: config.MatrixUsername,
				Color: "#0dbd8b",
			},
		})),
		emohtml.HStack(ctx, makeLinks([]link{
			{
				Name:  "mastodon",
				URL:   config.MastodonURL,
				Title: config.MastodonUsername,
				Color: "#6364ff",
			},
			{
				Name:  "github",
				URL:   config.GitHubURL,
				Title: config.GitHubUsername,
				Color: "#333",
			},
		})),
	)

	workedOn := emohtml.VStack(ctx,
		emohtml.StackSCSS(`
			align-items: start;
		`),
		makeLinks([]link{
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
				Icon:  "/icons/emoji/shark.svg",
			},
			{
				Name:  "baltimare leaderboard",
				Color: "#689F38",
				URL:   "https://baltimare.hotmilk.space",
				Icon:  "/icons/happy-anonfilly.png",
			},
			{
				Name:  "melon prime ds",
				Color: "#dd2e44",
				URL:   config.GitHubURL + "/melonPrimeDS",
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
				URL:   config.GitHubURL + "/dots",
				Icon:  "/icons/arch.svg",
				Muted: true,
			},
		}),
	)

	return Group{
		H2(Text("software engineer")),
		H2(Text("game developer")),
		H2(Text("server admin")),
		Br(),
		social,
		Br(),
		H2(Text("worked on")),
		Br(),
		workedOn,
		// Br(),
		// P(Text("may revert back to a new")),
		// P(Text("variant of the old page")),
	}
}
