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
	buttonWidth  = 88
	buttonHeight = 31
	renderScale  = 8

	// lower fps feels nicer for an 88x31
	fps    = 25 // 50 is max
	length = 4  // seconds
)

var (
	// should be less than render scale
	scales = []uint{1, 2}

	htmlPath    string
	localPath   string
	browserPool rod.Pool[rod.Browser]
	progress    *progressbar.ProgressBar
)

func createBrowser() (*rod.Browser, error) {
	return rod.New().ControlURL(
		launcher.New().Headless(true).MustLaunch(),
	).MustConnect(), nil
}

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
	browserPool = rod.NewBrowserPool(runtime.NumCPU())

	_, goFilename, _, _ := runtime.Caller(0)
	localPath = filepath.Dir(filepath.Clean(goFilename))

	htmlPath = filepath.Join(localPath, "index.html")

	wg := sync.WaitGroup{}

	// gifs are max 50 fps
	totalFrames := fps * length

	fmt.Println("rendering frames")
	progress = progressbar.Default(int64(totalFrames))

	for _, scale := range scales {
		os.Mkdir(getFramesDir(scale), 0755)
		os.Mkdir(getFramesDir(scale), 0755)
	}

	for i := range totalFrames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doFrame(i, scales)
		}()
	}

	wg.Wait()

	// gifski

	for _, scale := range scales {
		filename := "maki.gif"
		if scale > 1 {
			filename = fmt.Sprintf("maki@%dx.gif", scale)
		}

		outputGifFilePath := filepath.Join(
			localPath, "../../src/public/webring/", filename,
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
