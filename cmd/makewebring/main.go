package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/chromedp/chromedp"
	"github.com/disintegration/imaging"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/semaphore"
	"maki.cafe/cmd"
)

const (
	buttonWidth  = 88
	buttonHeight = 31
	renderScale  = 8

	// lower fps feels nicer for an 88x31
	fps    = 20 // 50 is max
	length = 2  // seconds
)

var (
	// should be less than render scale
	scales = []uint{1, 2}

	htmlPath  string
	localPath string
	chromeCtx context.Context
	progress  *progressbar.ProgressBar
)

func getFramesDir(scale uint) string {
	if scale > 1 {
		return filepath.Join(localPath, fmt.Sprintf("frames@%dx", scale))
	} else {
		return filepath.Join(localPath, "frames")
	}
}

func getFrameFilePath(i int, scale uint) string {
	return filepath.Join(getFramesDir(scale), fmt.Sprintf("%04d", i)+".png")
}

func doFrame(i int, scales []uint) {
	ctx, cancel := chromedp.NewContext(chromeCtx)
	defer cancel()

	var screenshotData []byte

	chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("file://" + htmlPath + "?go"),
		chromedp.EmulateViewport(buttonWidth, buttonHeight,
			chromedp.EmulateScale(renderScale),
		),
		chromedp.Evaluate(fmt.Sprintf("updateFrame(%d)", i), nil),
		chromedp.FullScreenshot(&screenshotData, 100),
	})

	frame, err := imaging.Decode(bytes.NewReader(screenshotData))
	if err != nil {
		panic(err)
	}

	for _, scale := range scales {
		scaled := imaging.Resize(
			frame, buttonWidth*int(scale), buttonHeight*int(scale),
			imaging.Lanczos,
		)

		file, err := os.OpenFile(
			getFrameFilePath(i, scale), os.O_RDWR|os.O_CREATE, 0644,
		)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = imaging.Encode(file, scaled, imaging.PNG)
		if err != nil {
			panic(err)
		}
	}

	progress.Add(1)
}

func main() {
	var cancel context.CancelFunc
	chromeCtx, cancel = chromedp.NewExecAllocator(
		context.Background(),
		chromedp.Headless,
	)
	defer cancel()

	_, goFilename, _, _ := runtime.Caller(0)
	localPath = filepath.Dir(filepath.Clean(goFilename))

	htmlPath = filepath.Join(localPath, "index.html")

	// gifs are max 50 fps
	totalFrames := fps * length

	fmt.Println("rendering frames")
	progress = progressbar.Default(int64(totalFrames))

	for _, scale := range scales {
		os.Mkdir(getFramesDir(scale), 0755)
	}

	{
		workers := int64(runtime.NumCPU())
		sem := semaphore.NewWeighted(workers)
		ctx := context.Background()

		for i := range totalFrames {
			sem.Acquire(ctx, 1)
			go func() {
				defer sem.Release(1)
				doFrame(i, scales)
			}()
		}

		sem.Acquire(ctx, workers)
	}

	// gifski

	for _, scale := range scales {
		filename := "maki.gif"
		if scale > 1 {
			filename = fmt.Sprintf("maki@%dx.gif", scale)
		}

		outputGifFilePath := filepath.Join(
			cmd.GetRootDir(), "src/public/webring/", filename,
		)

		gifskArgs := []string{
			"-W", strconv.Itoa(buttonWidth * int(scale)),
			"-H", strconv.Itoa(buttonHeight * int(scale)),
			"-r", strconv.Itoa(fps), // fps
			"-Q", "90", // quality
			"-o", outputGifFilePath, // quality
		}

		for i := range totalFrames {
			gifskArgs = append(gifskArgs, getFrameFilePath(i, scale))
		}

		cmd := exec.Command("gifski", gifskArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}
