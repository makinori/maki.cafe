package page

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/makinori/foxlib/foxcss"
	"github.com/makinori/foxlib/foxhtml"
	"maki.cafe/src/component"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	overwatchFileNameRegexp = regexp.MustCompile(
		`([0-9]{2,4})-([0-9]{1,2})-([0-9]{1,2})-(.+?)\..+$`,
	)
)

func Overwatch(ctx context.Context) Group {
	// allFiles, err := public.FS.ReadDir("overwatch")
	allFiles, err := os.ReadDir("big/overwatch")
	if err != nil {
		return Group{
			H2(Text("failed to get highlights")),
			H2(Text("try again later :(")),
		}
	}

	page := Group{
		component.IconHeader(ctx,
			H2(Text("highlights from over the years")),
			"/icons/overwatch.svg",
		),
	}

	videoFilenames := []string{}
	for _, file := range allFiles {
		ext := path.Ext(file.Name())
		if file.IsDir() || (ext != ".mp4" && ext != ".webm") {
			continue
		}
		videoFilenames = append(videoFilenames, file.Name())
	}

	slices.SortFunc(videoFilenames, func(a string, b string) int {
		return strings.Compare(b, a)
	})

	showErr := func(message string) {
		slog.Error("overwatch page: " + message)
		page = append(page, H3(Text(message)))
	}

	videoClass := foxcss.Class(ctx, `
		width: 100%;
		border-radius: 8px;
	`)

	for _, filename := range videoFilenames {
		matches := overwatchFileNameRegexp.FindStringSubmatch(filename)
		if len(matches) == 0 {
			showErr("failed to parse: " + filename)
			continue
		}

		year, err := strconv.ParseInt(matches[1], 10, 32)
		if err != nil {
			showErr("failed to parse: " + filename)
			continue
		}

		month, err := strconv.ParseInt(matches[2], 10, 32)
		if err != nil {
			showErr("failed to parse: " + filename)
			continue
		}

		day, err := strconv.ParseInt(matches[3], 10, 32)
		if err != nil {
			showErr("failed to parse: " + filename)
			continue
		}

		text := strings.ReplaceAll(matches[4], "-", " ")

		date := fmt.Sprintf(
			"%s %d '%d", strings.ToLower(time.Month(month).String()[:3]),
			day, year-2000,
		)

		page = append(page,
			foxhtml.HStack(ctx,
				foxhtml.StackSCSS(`
					margin-top: 24px;
					margin-bottom: 8px;
				`),
				H2(Text(text)),
				Div(Style("flex-grow:1")),
				H3(Text(date)),
			),
			Video(
				Controls(),
				Class(videoClass),
				Preload("none"),
				Src("/overwatch/"+filename),
				Poster("/overwatch/"+strings.TrimSuffix(
					filename, path.Ext(filename),
				)+".jpg"),
			),
		)
	}

	page = append(page,
		foxhtml.VStack(ctx,
			foxhtml.StackSCSS(`
				width: 400px;
				align-items: center;
				margin-top: 96px;
			`),
			Img(
				Width("400"),
				Src("/images/kiri-donut.jpg"),
			),
			P(I(Text("i love my wife kiriko"))),
		),
	)

	return page
}
