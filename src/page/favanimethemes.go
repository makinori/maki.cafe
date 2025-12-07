package page

import (
	"context"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/makinori/foxlib/foxcss"
	"maki.cafe/src/util"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func FavAnimeThemes(ctx context.Context) Group {
	allFiles, err := os.ReadDir("big/anime-themes")
	if err != nil {
		return Group{
			H2(Text("failed to get anime themes")),
			H2(Text("try again later :(")),
		}
	}

	page := Group{
		H2(Text("favorite anime openings and endings")),
		Br(),
		P(
			Text("in no particular order. videos from "),
			A(Href("https://animethemes.moe"), Text("animethemes.moe")),
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

	util.Shuffle(sort.StringSlice(videoFilenames))

	videoClass := foxcss.Class(ctx, `
		width: 100%;
		border-radius: 8px;
	`)

	for _, filename := range videoFilenames {
		filenameNoExt := strings.TrimSuffix(
			filename, path.Ext(filename),
		)

		title := ""
		link := ""

		info, err := os.ReadFile("big/anime-themes/" + filenameNoExt + ".txt")
		if err == nil {
			lines := strings.Split(string(info), "\n")
			title = strings.ToLower(strings.TrimSpace(lines[0]))
			link = strings.TrimSpace(lines[1])
		}

		page = append(page,
			H3(
				Class(foxcss.Class(ctx, `
					margin-top: 24px;
					margin-bottom: 8px;
				`)),
				A(
					Class("plain"),
					Href(link),
					Text(title),
				),
			),
			Video(
				Controls(),
				Class(videoClass),
				Preload("none"),
				Src("/anime-themes/"+filename),
				Poster("/anime-themes/"+filenameNoExt+".jpg"),
			),
		)
	}

	return page
}
