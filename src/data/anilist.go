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

	fmt.Println(string(queryJson))

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
	Current            []aniListCurrentAnime   `json:"current"`
	Completed          []aniListCompletedAnime `json:"completed"`
	FavoriteAnime      []aniListAnime          `json:"favoriteAnime"`
	FavoriteCharacters []aniListCharacter      `json:"favoriteCharacters"`
}

const aniListMaxPerPage = 50

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
			PerPage:  12,
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

	return aniListResult{
		Current:            current.Data.Page.MediaList,
		Completed:          completed.Data.Page.MediaList,
		FavoriteAnime:      favorites.Data.User.Favorites.Anime.Nodes,
		FavoriteCharacters: favorites.Data.User.Favorites.Characters.Nodes,
	}, nil
}

var Anilist = cachedData[aniListResult]{
	Key:      "anilist",
	CronSpec: "0 0 * * *", // once a day
	retrieve: getAniList,
}
