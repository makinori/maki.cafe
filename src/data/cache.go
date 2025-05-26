package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/disintegration/imaging"
	"maki.cafe/src/spritesheet"
)

func setCache(key string, data any, expiresAt time.Time) error {
	os.Mkdir("cache", 0755)

	finalData := map[string]any{
		"expiresAt": expiresAt,
		"data":      data,
	}

	jsonBytes, err := json.Marshal(finalData)
	if err != nil {
		return err
	}

	err = os.WriteFile("cache/"+key+".json", jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func writeCachePublic(filename string, data []byte) error {
	filePath := "cache/public/" + filename

	os.MkdirAll(path.Dir(filePath), 0755)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getCache[T any](key string, output *T) error {
	bytes, err := os.ReadFile("cache/" + key + ".json")
	if err != nil {
		return err
	}

	var cacheData struct {
		ExpiresAt time.Time `json:"expiresAt"`
		Data      T         `json:"data"`
	}

	err = json.Unmarshal(bytes, &cacheData)
	if err != nil {
		return err
	}

	if time.Now().After(cacheData.ExpiresAt) {
		os.Remove("cache/" + key + ".json")
		return errors.New("cache data expired")
	}

	*output = cacheData.Data

	return nil
}

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
