package page

import (
	"time"

	"github.com/makinori/maki.cafe/src/config"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/render"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Anime() Group {
	css, _ := render.RenderSass(`
		.anime-grid {
			display: grid;
			grid-template-columns: repeat(5,1fr);
			grid-gap: 8px;
		
			.anime {
				padding: 0;
				display: flex;
				flex-direction: column;
				gap: 4px;
				background: transparent;

				> div {
					aspect-ratio: 23/32;
					width: 100%;
					background-size: cover;
					background-position: center;
					background-color: #222;
				}

				> p {
					font-size: 20px;
				}
			}
		}
	`)

	var animeNodes []Node

	for _, anime := range data.Anilist.Data.Data.Page.MediaList {
		completedAt := time.Date(
			anime.CompletedAt.Year,
			time.Month(anime.CompletedAt.Month),
			anime.CompletedAt.Day,
			0, 0, 0, 0, time.UTC,
		)

		animeNodes = append(animeNodes, A(
			Href(anime.Media.SiteURL),
			Class("anime"),
			// Img(
			// 	Style("width:100%;height:auto"),
			// 	Src(anime.Media.CoverImage.Large),
			// ),
			Div(
				Style("background-image:url(\""+
					anime.Media.CoverImage.Large+
					"\")"),
			),
			P(Text(util.ShortDate(completedAt))),
		))
	}

	return Group{
		StyleEl(Raw(css)),
		P(
			Text("see my "),
			A(
				Href(config.AniListURL),
				Text("anilist profile"),
			),
			Text(" for more"),
		),
		Br(),
		H1(Text("recently finished")),
		Br(),
		Div(
			Class("anime-grid"),
			Group(animeNodes),
		),
		Br(),
		A(
			Href("https://anilist.co/user/"+config.AniListUsername+"/animelist/Completed"),
			Text("see all recently finished"),
			Class("muted"),
		),
	}
}
