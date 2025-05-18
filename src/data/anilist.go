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

type anilistQueryVars struct {
	UserName string `json:"userName"`
}

type anilistQuery struct {
	Query     string           `json:"query"`
	Variables anilistQueryVars `json:"variables"`
}

type anilistResult struct {
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

func getAnilist() (anilistResult, error) {
	aniListQueryJson, err := json.Marshal(anilistQuery{
		Query: aniListQuery,
		Variables: anilistQueryVars{
			UserName: config.AniListUsername,
		},
	})

	if err != nil {
		return anilistResult{}, err
	}

	req, err := http.NewRequest(
		"POST", "https://graphql.anilist.co",
		bytes.NewBuffer(aniListQueryJson),
	)

	if err != nil {
		return anilistResult{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return anilistResult{}, err
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return anilistResult{}, err
	}

	var result anilistResult
	err = json.Unmarshal(resBytes, &result)
	if err != nil {
		return anilistResult{}, err
	}

	if len(result.Errors) > 0 {
		return anilistResult{}, fmt.Errorf("%s", result.Errors)
	}

	return result, nil
}

var Anilist = cachedData[anilistResult]{
	Key:      "anilist",
	CronSpec: "0 0 * * *", // once a day
	retrieve: getAnilist,
}
