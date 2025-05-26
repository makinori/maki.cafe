package page

import (
	"context"
	"fmt"
	"time"

	"github.com/makinori/maki.cafe/src/component"
	"github.com/makinori/maki.cafe/src/config"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TODO: refactor name to favanime

func animeTitle(title data.AniListTitle) string {
	if title.English != "" {
		return title.English
	}
	return title.Romaji
}

func FavAnime(ctx context.Context) Group {
	anilistData := data.Anilist.Data

	var current, completed, favoriteAnime, favoriteCharacters Group

	for i, anime := range anilistData.Current {
		current = append(current, component.SpriteSheetGridItem(
			animeTitle(anime.Media.Title), anime.Media.SiteURL,
			anilistData.CurrentImage.Positions[i],
			P(Text(
				fmt.Sprintf("%d/%d", anime.Progress, anime.Media.Episodes),
			)),
		))
	}

	for i, anime := range anilistData.Completed {
		completedAt := time.Date(
			anime.CompletedAt.Year,
			time.Month(anime.CompletedAt.Month),
			anime.CompletedAt.Day,
			0, 0, 0, 0, time.UTC,
		)

		completed = append(completed, component.SpriteSheetGridItem(
			animeTitle(anime.Media.Title), anime.Media.SiteURL,
			anilistData.CompletedImage.Positions[i],
			P(Text(util.ShortDateWithYear(completedAt))),
		))
	}

	for i, anime := range anilistData.FavoriteAnime {
		favoriteAnime = append(favoriteAnime, component.SpriteSheetGridItem(
			animeTitle(anime.Title), anime.SiteURL,
			anilistData.FavoriteAnimeImage.Positions[i],
		))
	}

	for i, character := range anilistData.FavoriteCharacters {
		favoriteCharacters = append(favoriteCharacters, component.SpriteSheetGridItem(
			character.Name.UserPreferred, character.SiteURL,
			anilistData.FavoriteCharactersImage.Positions[i],
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
		component.SpriteSheetGrid(ctx,
			anilistData.CurrentImage.ImageURL,
			anilistData.CurrentImage.Size,
			data.AniListRatio, data.AniListGridWidthLarge, current,
		),
		Br(),
		H1(Text("recently finished")),
		Br(),
		component.SpriteSheetGrid(ctx,
			anilistData.CompletedImage.ImageURL,
			anilistData.CompletedImage.Size,
			data.AniListRatio, data.AniListGridWidthLarge, completed,
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
		component.SpriteSheetGrid(ctx,
			anilistData.FavoriteAnimeImage.ImageURL,
			anilistData.FavoriteAnimeImage.Size,
			data.AniListRatio, data.AniListGridWidthSmall, favoriteAnime,
		),
		Br(),
		H1(Text("favorite characters")),
		Br(),
		component.SpriteSheetGrid(ctx,
			anilistData.FavoriteCharactersImage.ImageURL,
			anilistData.FavoriteCharactersImage.Size,
			data.AniListRatio, data.AniListGridWidthSmall, favoriteCharacters,
		),
		// Br()
		// TODO: https://github.com/makinori/anilist-spinner
	}
}
