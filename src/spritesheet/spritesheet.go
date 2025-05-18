package spritesheet

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"math"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
	"golang.org/x/sync/semaphore"
)

func cssPercentageNumbers(x float64, y float64) string {
	return fmt.Sprintf(
		"%s%% %s%%",
		strings.ReplaceAll(fmt.Sprintf("%.2f", x*100), ".00", ""),
		strings.ReplaceAll(fmt.Sprintf("%.2f", y*100), ".00", ""),
	)
}

type SpriteSheetCSS struct {
	Size      string   `json:"size"`
	Positions []string `json:"positions"`
}

func Generate(
	imageWidth int, imageHeight int, imagePadding int,
	sheetWidth int, sheetHeight int,
	images []image.Image,
) (image.Image, SpriteSheetCSS, error) {
	if len(images) > sheetWidth*sheetHeight {
		return nil, SpriteSheetCSS{}, fmt.Errorf(
			"too many images. only %d spaces", sheetWidth*sheetHeight,
		)
	}

	spriteSheetWidth := imageWidth*sheetWidth + imagePadding*(sheetWidth-1)
	spriteSheetHeight := imageHeight*sheetHeight + imagePadding*(sheetHeight-1)

	css := SpriteSheetCSS{
		Size: cssPercentageNumbers(
			float64(spriteSheetWidth)/float64(imageWidth),
			float64(spriteSheetHeight)/float64(imageHeight),
		),
	}

	var finalMutex sync.Mutex
	final := imaging.New(spriteSheetWidth, spriteSheetHeight, color.Black)

	var concurrency int64 = int64(runtime.NumCPU())
	ctx := context.Background()
	sem := semaphore.NewWeighted(concurrency)

	for i, img := range images {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			return nil, SpriteSheetCSS{}, err
		}

		x := i % sheetWidth
		y := int(math.Floor(float64(i) / float64(sheetWidth)))

		css.Positions = append(css.Positions,
			cssPercentageNumbers(
				float64(x)/math.Max(1, float64(sheetWidth-1)),
				float64(y)/math.Max(1, float64(sheetHeight-1)),
			),
		)

		go func() {
			defer sem.Release(1)

			img = imaging.Fill(
				img, imageWidth, imageHeight, imaging.Center, imaging.Lanczos,
			)

			finalMutex.Lock()
			final = imaging.Paste(final, img, image.Point{
				X: x*imageWidth + x*imagePadding,
				Y: y*imageHeight + y*imagePadding,
			})
			finalMutex.Unlock()
		}()
	}

	err := sem.Acquire(ctx, concurrency)
	if err != nil {
		return nil, SpriteSheetCSS{}, err
	}

	return final, css, nil
}

func GenerateFromURLs(
	imageWidth int, imageHeight int, imagePadding int,
	sheetWidth int, sheetHeight int,
	urlsOrPaths []string,
) (image.Image, SpriteSheetCSS, error) {
	images := make([]image.Image, len(urlsOrPaths))
	var imagesMutex sync.Mutex

	var concurrency int64 = 8
	ctx := context.Background()
	sem := semaphore.NewWeighted(concurrency)

	var errs []error
	var errMutex sync.Mutex

	for i, url := range urlsOrPaths {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			return nil, SpriteSheetCSS{}, err
		}

		go func() {
			defer sem.Release(1)

			var image image.Image

			if strings.HasPrefix(url, "http") {
				res, err := http.Get(url)
				if err != nil {
					errMutex.Lock()
					errs = append(errs, err)
					errMutex.Unlock()
				}
				defer res.Body.Close()

				image, err = imaging.Decode(res.Body)
				if err != nil {
					errMutex.Lock()
					errs = append(errs, err)
					errMutex.Unlock()
					return
				}
			} else {
				data, err := os.ReadFile(url)
				if err != nil {
					errMutex.Lock()
					errs = append(errs, err)
					errMutex.Unlock()
				}

				image, err = imaging.Decode(bytes.NewReader(data))
				if err != nil {
					errMutex.Lock()
					errs = append(errs, err)
					errMutex.Unlock()
					return
				}
			}

			imagesMutex.Lock()
			images[i] = image
			imagesMutex.Unlock()
		}()
	}

	err := sem.Acquire(ctx, concurrency)
	if err != nil {
		return nil, SpriteSheetCSS{}, err
	}

	if len(errs) > 0 {
		return nil, SpriteSheetCSS{}, fmt.Errorf("%v", errs)
	}

	return Generate(
		imageWidth, imageHeight, imagePadding,
		sheetWidth, sheetHeight, images,
	)
}

func GenerateFromURLsGuessHeight(
	imageWidth int, imageHeight int, imagePadding int,
	sheetWidth int, urlsOrPaths []string,
) (image.Image, SpriteSheetCSS, error) {
	sheetHeight := int(
		math.Ceil(float64(len(urlsOrPaths)) / float64(sheetWidth)),
	)
	return GenerateFromURLs(
		imageWidth, imageHeight, imagePadding,
		sheetWidth, sheetHeight, urlsOrPaths,
	)
}
