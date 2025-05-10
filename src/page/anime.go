package page

import (
	"fmt"
	"time"

	"github.com/makinori/maki.cafe/src/common"
	"github.com/makinori/maki.cafe/src/page/anime"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// all very wip. will add caching, css-in-go and such

func Anime() Group {
	anilist, err := anime.GetAnilist()
	if err != nil {
		fmt.Println(err)
	}

	var animeNodes []Node

	for _, anime := range anilist.Data.Page.MediaList {
		completedAt := time.Date(
			anime.CompletedAt.Year,
			time.Month(anime.CompletedAt.Month),
			anime.CompletedAt.Day,
			0, 0, 0, 0, time.UTC,
		)

		animeNodes = append(animeNodes, A(
			Href(anime.Media.SiteURL),
			Style("padding:0;display:flex;flex-direction:column;gap:4px;background:transparent"),
			// Img(
			// 	Style("width:100%;height:auto"),
			// 	Src(anime.Media.CoverImage.Large),
			// ),
			Div(
				Style("aspect-ratio:23/32;width:100%;background-size:cover;background-position:center;background-color:#222;background-image:url(\""+anime.Media.CoverImage.Large+"\")"),
			),

			P(
				Style("font-size:20px"),
				Text(util.ShortDate(completedAt)),
			),
		))
	}

	return Group{
		P(
			Text("see my "),
			A(
				Href(common.AniListURL),
				Text("anilist profile"),
			),
			Text(" for more"),
		),
		Br(),
		H1(Text("recently finished")),
		Br(),
		Div(
			Style("display:grid;grid-template-columns:repeat(5,1fr);grid-gap:8px;"),
			Group(animeNodes),
		),
		Br(),
		A(
			Href("https://anilist.co/user/"+common.AniListUsername+"/animelist/Completed"),
			Text("see all recently finished"),
			Class("muted"),
		),
	}
}
