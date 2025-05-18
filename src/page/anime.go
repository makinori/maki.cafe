package page

import (
	"context"
	"fmt"
	"time"

	"github.com/makinori/maki.cafe/src/config"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/render"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func makeAnimeNode(
	ctx context.Context, name string, href string, image string,
	nodes ...Node,
) Node {
	return A(
		Title(name),
		Href(href),
		Class(render.SCSS(ctx, `
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
					font-size: 18px;
				}
			`)),
		// Img(
		// 	Style("width:100%;height:auto"),
		// 	Src(anime.Media.CoverImage.Large),
		// ),
		Div(
			Style(fmt.Sprintf(
				`background-image:url("%s")`, image,
			)),
		),
		Group(nodes),
	)
}

func animeTitle(title data.AniListTitle) string {
	if title.English != "" {
		return title.English
	}
	return title.Romaji
}

func Anime(ctx context.Context) Group {
	var current, completed, favoriteAnime, favoriteCharacters Group

	for _, anime := range data.Anilist.Data.Current {
		current = append(current, makeAnimeNode(
			ctx, animeTitle(anime.Media.Title),
			anime.Media.SiteURL, anime.Media.CoverImage.Large,
			P(Text(
				fmt.Sprintf("%d/%d", anime.Progress, anime.Media.Episodes),
			)),
		))
	}

	for _, anime := range data.Anilist.Data.Completed {
		completedAt := time.Date(
			anime.CompletedAt.Year,
			time.Month(anime.CompletedAt.Month),
			anime.CompletedAt.Day,
			0, 0, 0, 0, time.UTC,
		)

		completed = append(completed, makeAnimeNode(
			ctx, animeTitle(anime.Media.Title),
			anime.Media.SiteURL, anime.Media.CoverImage.Large,
			P(Text(util.ShortDateWithYear(completedAt))),
		))
	}

	for _, anime := range data.Anilist.Data.FavoriteAnime {
		favoriteAnime = append(favoriteAnime, makeAnimeNode(
			ctx, animeTitle(anime.Title),
			anime.SiteURL, anime.CoverImage.Large,
		))
	}

	for _, anime := range data.Anilist.Data.FavoriteCharacters {
		favoriteCharacters = append(favoriteCharacters, makeAnimeNode(
			ctx, anime.Name.UserPreferred,
			anime.SiteURL, anime.Image.Large,
		))
	}

	return Group{
		P(
			Text("see my "),
			A(
				Href(config.AniListURL),
				Text("anilist profile"),
			),
			Text(" for more"),
		),
		Br(),
		H1(Text("currently watching")),
		Br(),
		Div(
			Class(render.SCSS(ctx, `
				display: grid;
				grid-template-columns: repeat(6,1fr);
				grid-gap: 8px;
			`)),
			current,
		),
		Br(),
		H1(Text("recently finished")),
		Br(),
		Div(
			Class(render.SCSS(ctx, `
				display: grid;
				grid-template-columns: repeat(6,1fr);
				grid-gap: 8px;
			`)),
			completed,
		),
		// Br(),
		// A(
		// 	Href("https://anilist.co/user/"+config.AniListUsername+"/animelist/Completed"),
		// 	Text("see all finished"),
		// 	Class("muted"),
		// ),
		// Br(),
		Br(),
		H1(Text("favorite anime")),
		Br(),
		Div(
			Class(render.SCSS(ctx, `
				display: grid;
				grid-template-columns: repeat(8,1fr);
				grid-gap: 8px;
			`)),
			favoriteAnime,
		),
		Br(),
		H1(Text("favorite characters")),
		Br(),
		Div(
			Class(render.SCSS(ctx, `
				display: grid;
				grid-template-columns: repeat(8,1fr);
				grid-gap: 8px;
			`)),
			favoriteCharacters,
		),
		// Br()
		// TODO: https://github.com/makinori/anilist-spinner
	}
}
