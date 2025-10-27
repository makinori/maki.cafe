package data

import (
	"bytes"
	"fmt"
	"time"

	"github.com/disintegration/imaging"
	"maki.cafe/src/spritesheet"
)

type CachedSpriteSheet struct {
	ImageURL  string   `json:"imageURL"`
	Size      string   `json:"size"`
	Positions []string `json:"position"`
}

func makeCachedSpriteSheet[T any](
	name string, array *[]T, getImageURL func(el *T) string,
	imageWidth int, imageHeight int, imagePadding int, sheetWidth int,
) (CachedSpriteSheet, error) {
	var imageURLs []string
	for _, el := range *array {
		imageURLs = append(imageURLs, getImageURL(&el))
	}

	image, css, err := spritesheet.GenerateFromURLsGuessHeight(
		imageWidth, imageHeight, imagePadding, sheetWidth, imageURLs,
	)

	if err != nil {
		return CachedSpriteSheet{}, err
	}

	jpg := bytes.NewBuffer(nil)
	err = imaging.Encode(jpg, image, imaging.JPEG, imaging.JPEGQuality(90))

	if err != nil {
		return CachedSpriteSheet{}, err
	}

	filePath := name + ".jpg"
	err = writeCachePublic(filePath, jpg.Bytes())
	if err != nil {
		return CachedSpriteSheet{}, err
	}

	// is ?<time> necessary? if the images dont change, it'll still redownload

	return CachedSpriteSheet{
		ImageURL:  fmt.Sprintf("/cache/%s?%d", filePath, time.Now().Unix()),
		Size:      css.Size,
		Positions: css.Positions,
	}, nil
}
