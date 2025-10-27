package data

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/makinori/goemo/emocache"
	"maki.cafe/src/config"
)

type squirrelImage struct {
	Link      string    `json:"link"`
	Date      time.Time `json:"date"`
	Thumbnail string    `json:"thumbnail"`
}

type rssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		// Title string `xml:"title"`
		Items []struct {
			Link         string `xml:"link"`
			PubDate      string `xml:"pubDate"`
			MediaContent struct {
				URL string `xml:"url,attr"`
			} `xml:"content"`
		} `xml:"item"`
	} `xml:"channel"`
}

func getSquirrels() ([]squirrelImage, error) {
	res, err := http.Get(config.SquirrelsURL + ".rss")
	if err != nil {
		return []squirrelImage{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []squirrelImage{}, err
	}

	var rssFeed rssFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return []squirrelImage{}, err
	}

	var squirrelImages []squirrelImage

	for _, item := range rssFeed.Channel.Items {
		thumbnailURL := strings.ReplaceAll(
			item.MediaContent.URL, "/original/", "/small/",
		)

		date, _ := time.Parse(time.RFC1123Z, item.PubDate)

		squirrelImages = append(squirrelImages, squirrelImage{
			Link:      item.Link,
			Date:      date,
			Thumbnail: thumbnailURL,
		})
	}

	return squirrelImages, nil
}

var Squirrels = emocache.Data[[]squirrelImage]{
	Key:      "squirrels",
	CronSpec: "0 0 * * *", // once a day
	Retrieve: getSquirrels,
}
