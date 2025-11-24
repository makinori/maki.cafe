package page

import (
	"context"

	"github.com/makinori/foxlib/foxhtml"
	"maki.cafe/src/config"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type link struct {
	Name       string
	Color      string
	URL        string
	Icon       string
	IconBigger bool
	Muted      bool
	Break      bool
	Title      string
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
			imgParams := []Node{Src(item.Icon)}
			if item.IconBigger {
				imgParams = append(imgParams, Class("bigger"))
			}
			params = append(params, Img(imgParams...))
		}

		params = append(params, Text(item.Name))

		output = append(output, A(params))
	}

	return output
}

func Index(ctx context.Context) Group {
	// TODO: add icons?

	social := foxhtml.VStack(ctx,
		foxhtml.HStack(ctx, makeLinks([]link{
			{
				Name:  "email",
				URL:   "/email",
				Title: config.Email,
				// Color: "#333",
				Icon: "/icons/fa/envelope.svg",
			},
			{
				Name:  "xmpp",
				URL:   "/xmpp",
				Title: config.XMPP,
				Color: "#227ee1",
				Icon:  "/icons/xmpp.svg",
			},
			{
				Name:  "tox",
				URL:   config.ToxURI,
				Title: config.Tox,
				Color: "#ff8f00",
				Icon:  "/icons/tox.svg",
			},
			{
				Name:  "matrix",
				URL:   config.MatrixURL,
				Title: config.MatrixUsername,
				Color: "#0dbd8b",
				Icon:  "/icons/element.svg",
			},
		})),
		foxhtml.HStack(ctx, makeLinks([]link{
			{
				Name:  "mastodon",
				URL:   config.MastodonURL,
				Title: config.MastodonUsername,
				Color: "#6364ff",
				Icon:  "/icons/fa/mastodon.svg",
			},
			{
				Name:  "github",
				URL:   config.GitHubURL,
				Title: config.GitHubUsername,
				Color: "#333",
				Icon:  "/icons/fa/github.svg",
			},
		})),
		foxhtml.HStack(ctx, makeLinks([]link{
			{
				Name:  "second life",
				URL:   config.SecondLifeURL,
				Title: config.SecondLifeName,
				Color: "#00bfff",
				Icon:  "/icons/second-life.svg",
			},
		})),
	)

	workedOn := foxhtml.VStack(ctx,
		foxhtml.StackSCSS(`
		align-items: start;
		`),
		makeLinks([]link{
			{
				Name:       "tivoli cloud vr",
				Color:      "#e91e63",
				URL:        "https://github.com/tivolicloud",
				Icon:       "/icons/tivoli.svg",
				IconBigger: true,
			},
			{
				Name:       "blahaj quest",
				Color:      "#3c8ea7",
				URL:        "https://blahaj.quest",
				Icon:       "/icons/emoji/shark.svg",
				IconBigger: true,
			},
			{
				Name:       "baltimare leaderboard",
				Color:      "#689F38",
				URL:        "https://baltimare.hotmilk.space",
				Icon:       "/icons/happy-anonfilly.png",
				IconBigger: true,
			},
			{
				Name:       "melon prime ds",
				Color:      "#dd2e44",
				URL:        config.GitHubURL + "/melonPrimeDS",
				Icon:       "/icons/metroid.png",
				IconBigger: true,
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
		H2(
			Text("write software, make games"),
			Br(),
			Text("and run servers"),
		),
		H3(
			Style("margin-top: 16px; margin-bottom: 4px"),
			Text("also a cute fox girl"),
			Br(),
			Text("she/they"),
			Img(Src("/icons/trans-heart.svg"), Class("icon")),
		),
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
