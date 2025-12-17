package data

import (
	"context"
	"fmt"

	"github.com/hasura/go-graphql-client"
	"github.com/makinori/foxlib/foxcache"
	"maki.cafe/src/config"
)

// tried to combine paginated queries but if we request page > max pages
// it just returns the last page, which is no good.

// https://studio.apollographql.com/sandbox/explorer
// https://graphql.anilist.co

// https://docs.anilist.co/guide/rate-limiting
// 90 req/min we should absolutely be fine

// TODO: we're gonna run out of spritesheet space if we keep favoriting lol

type aniListTitle struct {
	English string `graphql:"english"`
	Romaji  string `graphql:"romaji"`
}

func (title *aniListTitle) String() string {
	if title.English != "" {
		return title.English
	}
	return title.Romaji
}

type aniListNextAiringEpisode struct {
	Episode int `graphql:"episode"`
}

type aniListAnime struct {
	SiteURL    string `graphql:"siteUrl"`
	CoverImage struct {
		Large string `graphql:"large"`
	} `graphql:"coverImage"`
	Title             aniListTitle             `graphql:"title"`
	Episodes          int                      `graphql:"episodes"`
	NextAiringEpisode aniListNextAiringEpisode `graphql:"nextAiringEpisode"`
}

type aniListCurrentAnime struct {
	Progress int          `graphql:"progress"`
	Media    aniListAnime `graphql:"media"`
}

type aniListPageInfo struct {
	HasNextPage bool `graphql:"hasNextPage"`
}

type aniListCurrentQuery struct {
	Page struct {
		MediaList []aniListCurrentAnime `graphql:"mediaList(userName: $username, status_in: [CURRENT], type: ANIME, sort: [UPDATED_TIME_DESC])"`
		// PageInfo  aniListPageInfo       `graphql:"pageInfo"`
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
		// PageInfo  aniListPageInfo         `graphql:"pageInfo"`
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

type aniListFavoriteAnimeQuery struct {
	User struct {
		Favorites struct {
			Anime struct {
				Nodes    []aniListAnime  `graphql:"nodes"`
				PageInfo aniListPageInfo `graphql:"pageInfo"`
			} `graphql:"anime(page: $page)"`
		} `graphql:"favourites"`
	} `graphql:"User(name: $username)"`
}

type aniListFavoriteCharactersQuery struct {
	User struct {
		Favorites struct {
			Characters struct {
				Nodes    []aniListCharacter `graphql:"nodes"`
				PageInfo aniListPageInfo    `graphql:"pageInfo"`
			} `graphql:"characters(page: $page)"`
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

const aniListRatioWidth = 23
const aniListRatioHeight = 32

var AniListRatio = fmt.Sprintf("%d/%d", aniListRatioWidth, aniListRatioHeight)

var AniListGridWidthLarge = 6
var AniListGridWidthSmall = 8

func getAllPages[T any](
	client *graphql.Client, makeQuery func() *T,
	hasNextPage func(query *T) bool,
	merge func(query *T),
) error {
	page := 1
	more := true
	var query *T

	for more {
		query = makeQuery()
		err := client.Query(context.Background(), query, map[string]any{
			"username": config.AniListUsername,
			"page":     page,
		})
		if err != nil {
			return err
		}

		// fmt.Printf("got page %d\n", page)

		page++
		merge(query)
		more = hasNextPage(query)
	}

	return nil
}

func getAniList() (aniListResult, error) {
	client := graphql.NewClient("https://graphql.anilist.co", nil)

	result := aniListResult{}

	// only get fixed amount for currently watching and completed

	var current aniListCurrentQuery
	err := client.Query(context.Background(), &current, map[string]any{
		"username": config.AniListUsername,
		"page":     0,
		"perPage":  AniListGridWidthLarge * 2,
	})
	if err != nil {
		return aniListResult{}, err
	}
	result.Current = current.Page.MediaList

	var completed aniListCompletedQuery
	err = client.Query(context.Background(), &completed, map[string]any{
		"username": config.AniListUsername,
		"page":     0,
		"perPage":  AniListGridWidthLarge * 2,
	})
	if err != nil {
		return aniListResult{}, err
	}
	result.Completed = completed.Page.MediaList

	err = getAllPages(
		client,
		func() *aniListFavoriteAnimeQuery {
			return &aniListFavoriteAnimeQuery{}
		},
		func(query *aniListFavoriteAnimeQuery) bool {
			return query.User.Favorites.Anime.PageInfo.HasNextPage
		},
		func(query *aniListFavoriteAnimeQuery) {
			result.FavoriteAnime = append(result.FavoriteAnime,
				query.User.Favorites.Anime.Nodes...,
			)
		},
	)
	if err != nil {
		return aniListResult{}, err
	}

	err = getAllPages(
		client,
		func() *aniListFavoriteCharactersQuery {
			return &aniListFavoriteCharactersQuery{}
		},
		func(query *aniListFavoriteCharactersQuery) bool {
			return query.User.Favorites.Characters.PageInfo.HasNextPage
		},
		func(query *aniListFavoriteCharactersQuery) {
			result.FavoriteCharacters = append(result.FavoriteCharacters,
				query.User.Favorites.Characters.Nodes...,
			)
		},
	)
	if err != nil {
		return aniListResult{}, err
	}

	// make spritesheets

	imageWidth := aniListRatioWidth * 6
	imageHeight := aniListRatioHeight * 6
	imagePadding := 8

	result.CurrentImage, err = makeCachedSpriteSheet(
		"anilist/current", &result.Current,
		func(e *aniListCurrentAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.CompletedImage, err = makeCachedSpriteSheet(
		"anilist/completed", &result.Completed,
		func(e *aniListCompletedAnime) string { return e.Media.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthLarge,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteAnimeImage, err = makeCachedSpriteSheet(
		"anilist/favorite-anime", &result.FavoriteAnime,
		func(e *aniListAnime) string { return e.CoverImage.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthSmall,
	)
	if err != nil {
		return aniListResult{}, err
	}

	result.FavoriteCharactersImage, err = makeCachedSpriteSheet(
		"anilist/favorite-characters", &result.FavoriteCharacters,
		func(e *aniListCharacter) string { return e.Image.Large },
		imageWidth, imageHeight, imagePadding, AniListGridWidthSmall,
	)
	if err != nil {
		return aniListResult{}, err
	}

	return result, nil
}

var Anilist = foxcache.Data[aniListResult]{
	Key:      "anilist",
	CronSpec: "0 0 * * *", // once a day
	Retrieve: getAniList,
}
