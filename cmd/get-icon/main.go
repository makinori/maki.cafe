package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/svg"
)

func toCodePoint(unicodeSurrogates string, sep string) string {
	var r []string
	var p rune
	for _, c := range unicodeSurrogates {
		if p != 0 {
			r = append(r,
				fmt.Sprintf("%x", (0x10000+((p-0xd800)<<10)+(c-0xdc00))),
			)
			p = 0
		} else if 0xd800 <= c && c <= 0xdbff {
			p = c
		} else {
			r = append(r, fmt.Sprintf("%x", c))
		}
	}
	return strings.Join(r, sep)
}

func emojiURL(codepoint string, provider string) string {
	switch provider {
	case "noto":
		return "https://cdn.statically.io/gh/googlefonts/noto-emoji/main/svg/emoji_u" +
			strings.ReplaceAll(codepoint, "_fe0f", "") + ".svg"
	case "twemoji":
		return "https://cdnjs.cloudflare.com/ajax/libs/twemoji/14.0.2/svg/" +
			codepoint + ".svg"
	}
	return ""
}

func minifySVG(data []byte) ([]byte, error) {
	out := bytes.NewBuffer(nil)

	m := minify.New()
	m.Add("image/svg+xml", &svg.Minifier{})
	err := m.Minify("image/svg+xml", out, bytes.NewBuffer(data))
	if err != nil {
		return []byte{}, err
	}

	return out.Bytes(), nil
}

var usage = strings.TrimSpace(`
usage:
    emoji <name> <emoji> <provider>
    fa <name> <pack>
`)

func printExit1(a ...any) {
	fmt.Println(a...)
	os.Exit(1)
}

func getNextArgOrUsage(i *int) string {
	if *i >= len(os.Args) {
		printExit1(usage)
	}

	val := os.Args[*i]
	*i++

	if len(val) == 0 {
		printExit1(usage)
	}

	return val
}

func main() {
	argI := 1

	tool := getNextArgOrUsage(&argI)
	name := getNextArgOrUsage(&argI)

	var fileExt, url string

	switch tool {
	case "emoji":
		emoji := getNextArgOrUsage(&argI)
		provider := getNextArgOrUsage(&argI)
		codepoint := toCodePoint(emoji, "_")

		fileExt = "svg"
		url = emojiURL(codepoint, provider)

		if url == "" {
			printExit1("unknown emoji provider: " + provider)
		}

	case "fa":
		// solid, brands
		pack := getNextArgOrUsage(&argI)
		if pack == "" {
			printExit1("unknown fa pack: " + pack)
		}

		fileExt = "svg"
		url = "https://site-assets.fontawesome.com/releases/v6.7.2/svgs/" +
			pack + "/" + name + ".svg"

	default:
		printExit1(usage)
	}

	outputPath := "src/public/icons/" + tool + "/" + name + "." + fileExt

	// validation

	dirInfo, err := os.Stat(path.Dir(outputPath))
	if err != nil {
		printExit1(err)
	}
	if !dirInfo.IsDir() {
		printExit1(path.Dir(outputPath) + ": not a directory")
	}

	// download

	res, err := http.Get(url)
	if err != nil {
		printExit1(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		printExit1(err)
	}

	if res.StatusCode/100 != 2 {
		printExit1(string(data))
	}

	// optimize

	if fileExt == "svg" {
		sizeBefore := float64(len(data))
		data, err = minifySVG(data)
		if err != nil {
			printExit1(err)
		}

		fmt.Printf("optimized: %.1f%% of original\n", (float64(len(data))/sizeBefore)*100)
	}

	// save

	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		printExit1(err)
	}

	fmt.Println("saved to: " + outputPath)
}
