package page

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/makinori/maki.cafe/src/config"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/render"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TODO: refactor name to favanime and use grid component

func makeAnimeNode(
	ctx context.Context, name string, href string, image string,
	ss data.CachedSpriteSheet, index int,
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
				aspect-ratio: `+data.AniListRatio+`;
				width: 100%;
				background-color: #222;
			}

			> p {
				font-size: 18px;
			}
		`)),
		Div(
			Class(render.SCSS(ctx, `
				background-image: url("`+ss.ImageURL+`");
				background-size: `+ss.Size+`;
			`)),
			Style(fmt.Sprintf(
				`background-position: %s;`, ss.Positions[index],
			))),
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

	for i, anime := range data.Anilist.Data.Current {
		current = append(current, makeAnimeNode(
			ctx, animeTitle(anime.Media.Title),
			anime.Media.SiteURL, anime.Media.CoverImage.Large,
			data.Anilist.Data.CurrentImage, i,
			P(Text(
				fmt.Sprintf("%d/%d", anime.Progress, anime.Media.Episodes),
			)),
		))
	}

	for i, anime := range data.Anilist.Data.Completed {
		completedAt := time.Date(
			anime.CompletedAt.Year,
			time.Month(anime.CompletedAt.Month),
			anime.CompletedAt.Day,
			0, 0, 0, 0, time.UTC,
		)

		completed = append(completed, makeAnimeNode(
			ctx, animeTitle(anime.Media.Title),
			anime.Media.SiteURL, anime.Media.CoverImage.Large,
			data.Anilist.Data.CompletedImage, i,
			P(Text(util.ShortDateWithYear(completedAt))),
		))
	}

	for i, anime := range data.Anilist.Data.FavoriteAnime {
		favoriteAnime = append(favoriteAnime, makeAnimeNode(
			ctx, animeTitle(anime.Title),
			anime.SiteURL, anime.CoverImage.Large,
			data.Anilist.Data.FavoriteAnimeImage, i,
		))
	}

	for i, character := range data.Anilist.Data.FavoriteCharacters {
		favoriteCharacters = append(favoriteCharacters, makeAnimeNode(
			ctx, character.Name.UserPreferred,
			character.SiteURL, character.Image.Large,
			data.Anilist.Data.FavoriteCharactersImage, i,
		))
	}

	largeGrid := render.SCSS(ctx, `
		display: grid;
		grid-template-columns: repeat(`+strconv.Itoa(data.AniListGridWidthLarge)+`,1fr);
		grid-gap: 8px;
	`)

	smallGrid := render.SCSS(ctx, `
		display: grid;
		grid-template-columns: repeat(`+strconv.Itoa(data.AniListGridWidthSmall)+`,1fr);
		grid-gap: 8px;
	`)

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
		Div(Class(largeGrid), current),
		Br(),
		H1(Text("recently finished")),
		Br(),
		Div(Class(largeGrid), completed),
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
		Div(Class(smallGrid), favoriteAnime),
		Br(),
		H1(Text("favorite characters")),
		Br(),
		Div(Class(smallGrid), favoriteCharacters),
		// Br()
		// TODO: https://github.com/makinori/anilist-spinner
	}
}
