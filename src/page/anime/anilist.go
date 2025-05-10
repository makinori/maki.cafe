package anime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/makinori/maki.cafe/src/common"
)

// https://studio.apollographql.com/sandbox/explorer

var aniListQuery = `query ($userName: String) {
	Page(page: 0, perPage: 10) {
	  mediaList(
		userName: $userName
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
		  siteUrl
		  coverImage {
			large
		  }
		}
	  }
	}
  }`

type AnilistQueryVars struct {
	UserName string `json:"userName"`
}

type AnilistQuery struct {
	Query     string           `json:"query"`
	Variables AnilistQueryVars `json:"variables"`
}

type AnilistQueryResult struct {
	Errors []struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	} `json:"errors,omitempty"`
	Data struct {
		Page struct {
			MediaList []struct {
				CompletedAt struct {
					Day   int `json:"day"`
					Month int `json:"month"`
					Year  int `json:"year"`
				} `json:"completedAt"`
				Media struct {
					SiteURL    string `json:"siteUrl"`
					CoverImage struct {
						Large string `json:"large"`
					} `json:"coverImage"`
				} `json:"media"`
			} `json:"mediaList"`
		} `json:"Page"`
	} `json:"data"`
}

func GetAnilist() (AnilistQueryResult, error) {
	cacheBytes, err := os.ReadFile("anime.json")
	if err == nil {
		var result AnilistQueryResult
		err = json.Unmarshal(cacheBytes, &result)
		if err != nil {
			return AnilistQueryResult{}, err
		}

		return result, nil
	}

	aniListQueryJson, err := json.Marshal(AnilistQuery{
		Query: aniListQuery,
		Variables: AnilistQueryVars{
			UserName: common.AniListUsername,
		},
	})

	if err != nil {
		return AnilistQueryResult{}, err
	}

	req, err := http.NewRequest(
		"POST", "https://graphql.anilist.co",
		bytes.NewBuffer(aniListQueryJson),
	)

	if err != nil {
		return AnilistQueryResult{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return AnilistQueryResult{}, err
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return AnilistQueryResult{}, err
	}

	var result AnilistQueryResult
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return AnilistQueryResult{}, err
	}

	if len(result.Errors) > 0 {
		return AnilistQueryResult{}, fmt.Errorf("%s", result.Errors)
	}

	cacheBytes, err = json.Marshal(result)
	if err != nil {
		return AnilistQueryResult{}, err
	}

	err = os.WriteFile("anime.json", cacheBytes, 0644)
	if err != nil {
		return AnilistQueryResult{}, err
	}

	return result, nil
}
