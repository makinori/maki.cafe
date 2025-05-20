package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/schollz/progressbar/v3"
)

const (
	buttonWidth   = 88
	buttonHeight  = 31
	renderScale   = 8
	dontDownscale = false

	// lower fps feels nicer for an 88x31
	fps    = 25 // 50 is max
	length = 4  // seconds
)

var (
	htmlPath    string
	framesDir   string
	browserPool rod.Pool[rod.Browser]
	progress    *progressbar.ProgressBar
)

func createBrowser() (*rod.Browser, error) {
	return rod.New().ControlURL(
		launcher.New().Headless(true).MustLaunch(),
	).MustConnect(), nil
}

func getFrameFilePath(i int) string {
	return filepath.Join(framesDir, fmt.Sprintf("%04d", i)+".png")
}

func doFrame(i int) {
	browser, err := browserPool.Get(createBrowser)
	if err != nil {
		panic(err)
	}
	defer browserPool.Put(browser)

	page := browser.MustPage("file://" + htmlPath + "?go")
	defer page.Close()

	page.MustSetViewport(buttonWidth, buttonHeight, renderScale, false)
	page.MustWaitStable()
	page.MustEval("updateFrame", i)

	screenshotData, err := page.Screenshot(false, &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
	})
	if err != nil {
		panic(err)
	}

	// process image

	image, err := imaging.Decode(bytes.NewReader(screenshotData))
	if err != nil {
		panic(err)
	}

	// i think gifski uses lanczos. do it manually anyway
	if !dontDownscale {
		image = imaging.Resize(image, buttonWidth, buttonHeight, imaging.Lanczos)
	}

	// write

	file, err := os.OpenFile(getFrameFilePath(i), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = imaging.Encode(file, image, imaging.PNG)
	if err != nil {
		panic(err)
	}

	progress.Add(1)
}

func main() {
	browserPool = rod.NewBrowserPool(runtime.NumCPU())

	_, goFilename, _, _ := runtime.Caller(0)
	localPath := filepath.Dir(filepath.Clean(goFilename))

	htmlPath = filepath.Join(localPath, "index.html")
	framesDir = filepath.Join(localPath, "frames")

	wg := sync.WaitGroup{}

	// gifs are max 50 fps
	totalFrames := fps * length

	if dontDownscale {
		fmt.Println("warning: not downscaling")
	} else {

	}

	fmt.Println("rendering frames")
	progress = progressbar.Default(int64(totalFrames))

	for i := range totalFrames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doFrame(i)
		}()
	}

	wg.Wait()

	// gifski

	outputGifFilePath := filepath.Join(localPath, "../../src/public/webring/maki.gif")

	gifskArgs := []string{
		"-W", strconv.Itoa(buttonWidth),
		"-H", strconv.Itoa(buttonHeight),
		"-r", strconv.Itoa(fps), // fps
		"-Q", "90", // quality
		"-o", outputGifFilePath, // quality
	}

	for i := range totalFrames {
		gifskArgs = append(gifskArgs, getFrameFilePath(i))
	}

	cmd := exec.Command("gifski", gifskArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
