package page

import (
	"context"
	"strings"

	"github.com/makinori/foxlib/foxcss"
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
	SideText   string
}

func makeLinks(ctx context.Context, links []link) Group {
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

		if item.SideText == "" {
			output = append(output, A(params))
			continue
		}

		// side text only

		var sideTextParams Node
		if strings.HasPrefix(item.SideText, "(") {
			sideTextParams = Class(foxcss.Class(ctx, `
				opacity: 0.5;
			`))
		}

		output = append(output, foxhtml.HStack(ctx,
			foxhtml.StackSCSS(`
				align-items: center;
				gap: 12px;
			`),
			A(params),
			P(Text(item.SideText), sideTextParams),
		))
	}

	return output
}

func Index(ctx context.Context) Group {
	// TODO: add icons?

	reachMe := foxhtml.VStack(ctx,
		foxhtml.StackSCSS(`
			margin-bottom: 8px;
		`),
		foxhtml.HStack(ctx, makeLinks(ctx, []link{
			{
				Name:  "email",
				URL:   "/email",
				Title: config.Email,
				Icon:  "/icons/fa/envelope.svg",
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
		// foxhtml.HStack(ctx, makeLinks(ctx, []link{
		// 	{
		// 		Name:  "forgejo",
		// 		URL:   config.ForgejoURL,
		// 		Title: config.ForgejoDomain,
		// 		Color: "#ff5500",
		// 		Icon:  "/icons/forgejo.svg",
		// 	},
		// })),
		// foxhtml.HStack(ctx, makeLinks(ctx, []link{
		// 	{
		// 		Name:  "second life",
		// 		URL:   config.SecondLifeURL,
		// 		Title: config.SecondLifeName,
		// 		Color: "#00bfff",
		// 		Icon:  "/icons/second-life.svg",
		// 	},
		// })),
	)

	stackedLinks := foxhtml.StackSCSS(`
		align-items: start;
	`)

	workedOn := foxhtml.VStack(ctx,
		stackedLinks,
		makeLinks(ctx, []link{
			{
				Name:       "tivoli",
				Color:      "#e91e63",
				URL:        "https://github.com/tivolicloud",
				Icon:       "/icons/tivoli.svg",
				IconBigger: true,
				SideText:   "(archive)",
			},
			{
				Name:       "blahaj quest",
				Color:      "#3c8ea7",
				URL:        "https://blahaj.quest",
				Icon:       "/icons/emoji/shark.svg",
				IconBigger: true,
			},
			{
				Name:       "balti leaderboard",
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
		}),
	)

	findMore := foxhtml.HStack(ctx, makeLinks(ctx, []link{
		{
			Name:  "github",
			URL:   config.GitHubURL,
			Title: config.GitHubUsername,
			Color: "#333",
			Icon:  "/icons/fa/github.svg",
		},
		{
			Name:  "mastodon",
			URL:   config.MastodonURL,
			Title: config.MastodonUsername,
			Color: "#6364ff",
			Icon:  "/icons/fa/mastodon.svg",
		},
	}))

	goodiesFor := foxhtml.VStack(ctx,
		stackedLinks,
		makeLinks(ctx, []link{
			{
				Name:  "blender",
				Color: "#f4792b",
				URL:   "/dl/blender",
				Icon:  "/icons/blender.svg",
			},
		}),
	)

	// otherLinks := foxhtml.VStack(ctx,
	// 	stackedLinks,
	// 	makeLinks(ctx, []link{
	// 		// {
	// 		// 	Name:  "old page",
	// 		// 	Color: "#fff",
	// 		// 	URL:   "https://old.maki.cafe",
	// 		// 	Muted: true,
	// 		// },
	// 		{
	// 			Name:  "dots",
	// 			Color: "#fff",
	// 			URL:   config.GitHubURL + "/dots",
	// 			Icon:  "/icons/arch.svg",
	// 			Muted: true,
	// 		},
	// 	}),
	// )

	ircLink := func(name string, domain string) Node {
		return A(
			Class("plain"),
			Style("font-weight:600"),
			Href("https://"+domain),
			Text(name),
		)
	}

	return Group{
		H2(
			Text("write software, make games"),
			Br(),
			Text("and run servers"),
		),
		Br(),
		H2(Text("reach me")),
		Br(),
		reachMe,
		P(
			Text("or "),
			Code(Text("maki")),
			Text(" on "),
			ircLink("libera", "libera.chat"),
			Text(", "),
			ircLink("rizon", "rizon.net"),
			Text(" or "),
			ircLink("t420", "t420.net"),
		),
		Br(),
		H2(Text("worked on")),
		Br(),
		workedOn,
		Br(),
		H2(Text("find more")),
		Br(),
		findMore,
		Br(),
		H2(Text("goodies for")),
		Br(),
		goodiesFor,
		// Br(),
		// H2(Text("other")),
		// Br(),
		// otherLinks,
	}
}
