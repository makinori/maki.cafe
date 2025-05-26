package data

import (
	"context"
	"fmt"

	"github.com/hasura/go-graphql-client"
	"maki.cafe/src/config"
)

// https://studio.apollographql.com/sandbox/explorer
// https://graphql.anilist.co

type AniListTitle struct {
	English string `graphql:"english"`
	Romaji  string `graphql:"romaji"`
}

type aniListAnime struct {
	SiteURL    string `graphql:"siteUrl"`
	CoverImage struct {
		Large string `graphql:"large"`
	} `graphql:"coverImage"`
	Title    AniListTitle `graphql:"title"`
	Episodes int          `graphql:"episodes"`
}

type aniListCurrentAnime struct {
	Progress int          `graphql:"progress"`
	Media    aniListAnime `graphql:"media"`
}

type aniListCurrentQuery struct {
	Page struct {
		MediaList []aniListCurrentAnime `graphql:"mediaList(userName: $username, status_in: [CURRENT], type: ANIME, sort: [UPDATED_TIME_DESC])"`
	} `graphql:"Page(page: $page, perPage: $perPage)"`
}

type aniListCompletedAnime struct {
	CompletedAt struct {
		Day   int `graphql:"day"`
		Month int `graphql:"month"`
		Year  int `graphql:"year"`
	} `graphql:"completedAt"`
	Media aniListAnime `graphql:"media"`
}

type aniListCompletedQuery struct {
	Page struct {
		MediaList []aniListCompletedAnime `graphql:"mediaList(userName: $username, status_in: [COMPLETED], type: ANIME, sort: [FINISHED_ON_DESC])"`
	} `graphql:"Page(page: $page, perPage: $perPage)"`
}

type aniListCharacter struct {
	SiteURL string `graphql:"siteUrl"`
	Image   struct {
		Large string `graphql:"large"`
	} `graphql:"image"`
	Name struct {
		UserPreferred string `graphql:"userPreferred"`
	} `graphql:"name"`
}

type aniListFavoritesQuery struct {
	User struct {
		Favorites struct {
			Anime struct {
				Nodes []aniListAnime `graphql:"nodes"`
			} `graphql:"anime(page: $page, perPage: $perPage)"`
			Characters struct {
				Nodes []aniListCharacter `graphql:"nodes"`
			} `graphql:"characters(page: $page, perPage: $perPage)"`
		} `graphql:"favourites"`
	} `graphql:"User(name: $username)"`
}

// final struct

type aniListResult struct {
	Current                 []aniListCurrentAnime   `json:"current"`
	CurrentImage            CachedSpriteSheet       `json:"currentImage"`
	Completed               []aniListCompletedAnime `json:"completed"`
	CompletedImage          CachedSpriteSheet       `json:"completedImage"`
	FavoriteAnime           []aniListAnime          `json:"favoriteAnime"`
	FavoriteAnimeImage      CachedSpriteSheet       `json:"favoriteAnimeImage"`
	FavoriteCharacters      []aniListCharacter      `json:"favoriteCharacters"`
	FavoriteCharactersImage CachedSpriteSheet       `json:"favoriteCharactersImage"`
}

const aniListMaxPerPage = 50
const aniListRatioWidth = 23
const aniListRatioHeight = 32

var AniListRatio = fmt.Sprintf("%d/%d", aniListRatioWidth, aniListRatioHeight)

var AniListGridWidthLarge = 6
var AniListGridWidthSmall = 8

func getAniList() (aniListResult, error) {
	client := graphql.NewClient("https://graphql.anilist.co", nil)

	var current aniListCurrentQuery
	err := client.Query(context.Background(), &current, map[string]any{
		"username": config.AniListUsername,
		"page":     0,
		"perPage":  aniListMaxPerPage,
	})
	if err != nil {
		return aniListResult{}, err
	}

	var completed aniListCompletedQuery
	err = client.Query(context.Background(), &completed, map[string]any{
		"username": config.AniListUsername,
		"page":     0,
		"perPage":  AniListGridWidthLarge * 2,
	})
	if err != nil {
		return aniListResult{}, err
	}

	var favorites aniListFavoritesQuery
	err = client.Query(context.Background(), &favorites, map[string]any{
		"username": config.AniListUsername,
		"page":     0,
		"perPage":  aniListMaxPerPage,
	})
	if err != nil {
		return aniListResult{}, err
	}

	// make spritesheets

	result := aniListResult{
		Current:            current.Page.MediaList,
		Completed:          completed.Page.MediaList,
		FavoriteAnime:      favorites.User.Favorites.Anime.Nodes,
		FavoriteCharacters: favorites.User.Favorites.Characters.Nodes,
	}

	imageWidth := aniListRatioWidth * 6
	imageHeight := aniListRatioHeight * 6
	imagePadding := 8

	result.CurrentImage, err = makeCachedSpriteSheet(
		"anilist/current", &current.Page.MediaList,
		func(e *aniListCurrentAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.CompletedImage, err = makeCachedSpriteSheet(
		"anilist/completed", &completed.Page.MediaList,
		func(e *aniListCompletedAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteAnimeImage, err = makeCachedSpriteSheet(
		"anilist/favorite-anime", &favorites.User.Favorites.Anime.Nodes,
		func(e *aniListAnime) string { return e.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthSmall,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteCharactersImage, err = makeCachedSpriteSheet(
		"anilist/favorite-characters", &favorites.User.Favorites.Characters.Nodes,
		func(e *aniListCharacter) string { return e.Image.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthSmall,
	)
	if err != nil {
		return aniListResult{}, err
	}

	return result, nil
}

var Anilist = cachedData[aniListResult]{
	Key:      "anilist",
	CronSpec: "0 0 * * *", // once a day
	retrieve: getAniList,
}
