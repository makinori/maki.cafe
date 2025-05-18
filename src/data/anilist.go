package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/makinori/maki.cafe/src/config"
)

// https://studio.apollographql.com/sandbox/explorer
// https://graphql.anilist.co

var aniListAnimeSlice = `
siteUrl
coverImage {
	large
}
title {
	english
	romaji
}
`

type AniListTitle struct {
	English string `json:"english"`
	Romaji  string `json:"romaji"`
}

type aniListAnime struct {
	SiteURL    string `json:"siteUrl"`
	CoverImage struct {
		Large string `json:"large"`
	} `json:"coverImage"`
	Title    AniListTitle `json:"title"`
	Episodes int          `json:"episodes,omitempty"`
}

var aniListCurrentQuery = `
query ($username: String, $page: Int, $perPage: Int) {
	Page(page: $page, perPage: $perPage) {
		mediaList(
			userName: $username
			status_in: [CURRENT]
			type: ANIME
			sort: [STARTED_ON_DESC]
		) {
			progress
			media {
				` + aniListAnimeSlice + `
				episodes
			}
		}
	}
}
`

type aniListCurrentAnime struct {
	Progress int          `json:"progress"`
	Media    aniListAnime `json:"media"`
}

type aniListCurrentResult struct {
	Data struct {
		Page struct {
			MediaList []aniListCurrentAnime `json:"mediaList"`
		} `json:"Page"`
	} `json:"data"`
}

var aniListCompletedQuery = `
query ($username: String, $page: Int, $perPage: Int) {
	Page(page: $page, perPage: $perPage) {
		mediaList(
			userName: $username
			status_in: [COMPLETED]
			type: ANIME
			sort: [FINISHED_ON_DESC]
		) {
			completedAt {
				day
				month
				year
			}
			media {
				` + aniListAnimeSlice + `
			}
		}
	}
}
`

type aniListCompletedAnime struct {
	CompletedAt struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"completedAt"`
	Media aniListAnime `json:"media"`
}

type aniListCompletedResult struct {
	Data struct {
		Page struct {
			MediaList []aniListCompletedAnime `json:"mediaList"`
		} `json:"Page"`
	} `json:"data"`
}

var aniListCharacterQuery = `
siteUrl
image {
	large
}
name {
	userPreferred
}
`

type aniListCharacter struct {
	SiteURL string `json:"siteUrl"`
	Image   struct {
		Large string `json:"large"`
	} `json:"image"`
	Name struct {
		UserPreferred string `json:"userPreferred"`
	} `json:"name"`
}

var aniListFavoritesQuery = `
query ($username: String, $page: Int, $perPage: Int) {
	User(name: $username) {
		favourites {
			anime(page: $page, perPage: $perPage) {
				nodes {
` + aniListAnimeSlice + `
				}
			}
			characters(page: $page, perPage: $perPage) {
				nodes {
` + aniListCharacterQuery + `
				}
			}
		}
	}
}
`

type aniListFavoritesResult struct {
	Data struct {
		User struct {
			Favorites struct {
				Anime struct {
					Nodes []aniListAnime `json:"nodes"`
				} `json:"anime"`
				Characters struct {
					Nodes []aniListCharacter `json:"nodes"`
				} `json:"characters"`
			} `json:"favourites"`
		} `json:"User"`
	} `json:"data"`
}

// generic types

type aniListQueryVars struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PerPage  int    `json:"perPage"`
}

type aniListQuery struct {
	Query     string           `json:"query"`
	Variables aniListQueryVars `json:"variables"`
}

type aniListErrors struct {
	Errors []struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	} `json:"errors"`
}

func getQuery[T any](result *T, query aniListQuery) error {
	queryJson, err := json.Marshal(query)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST", "https://graphql.anilist.co",
		bytes.NewBuffer(queryJson),
	)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var errors aniListErrors
	err = json.Unmarshal(resBytes, &errors)
	if err != nil {
		return err
	}

	if len(errors.Errors) > 0 {
		return fmt.Errorf("%v", errors.Errors)
	}

	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return err
	}

	return nil
}

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
	var current aniListCurrentResult
	err := getQuery(&current, aniListQuery{
		Query: aniListCurrentQuery,
		Variables: aniListQueryVars{
			Username: config.AniListUsername,
			Page:     0,
			PerPage:  aniListMaxPerPage,
		},
	})

	if err != nil {
		return aniListResult{}, err
	}

	var completed aniListCompletedResult
	err = getQuery(&completed, aniListQuery{
		Query: aniListCompletedQuery,
		Variables: aniListQueryVars{
			Username: config.AniListUsername,
			Page:     0,
			PerPage:  AniListGridWidthLarge * 2,
		},
	})

	if err != nil {
		return aniListResult{}, err
	}

	var favorites aniListFavoritesResult
	err = getQuery(&favorites, aniListQuery{
		Query: aniListFavoritesQuery,
		Variables: aniListQueryVars{
			Username: config.AniListUsername,
			Page:     0,
			PerPage:  aniListMaxPerPage,
		},
	})

	if err != nil {
		return aniListResult{}, err
	}

	// make spritesheets

	result := aniListResult{
		Current:            current.Data.Page.MediaList,
		Completed:          completed.Data.Page.MediaList,
		FavoriteAnime:      favorites.Data.User.Favorites.Anime.Nodes,
		FavoriteCharacters: favorites.Data.User.Favorites.Characters.Nodes,
	}

	imageWidth := aniListRatioWidth * 6
	imageHeight := aniListRatioHeight * 6
	imagePadding := 8

	result.CurrentImage, err = makeCachedSpriteSheet(
		"anilist/current", &current.Data.Page.MediaList,
		func(e *aniListCurrentAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.CompletedImage, err = makeCachedSpriteSheet(
		"anilist/completed", &completed.Data.Page.MediaList,
		func(e *aniListCompletedAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteAnimeImage, err = makeCachedSpriteSheet(
		"anilist/favorite-anime", &favorites.Data.User.Favorites.Anime.Nodes,
		func(e *aniListAnime) string { return e.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthSmall,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteCharactersImage, err = makeCachedSpriteSheet(
		"anilist/favorite-characters", &favorites.Data.User.Favorites.Characters.Nodes,
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
