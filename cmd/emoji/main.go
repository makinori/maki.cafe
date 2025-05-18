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

func main() {
	if len(os.Args) < 4 {
		panic("usage: <emoji> <provider> <name>")
	}

	emoji := os.Args[1]
	provider := os.Args[2]
	name := os.Args[3]

	codepoint := toCodePoint(emoji, "_")
	url := emojiURL(codepoint, provider)
	outputPath := "src/public/emojis/" + name + ".svg"

	// validation

	if url == "" {
		panic("unknown provider: " + provider)
	}

	dirInfo, err := os.Stat(path.Dir(outputPath))
	if err != nil {
		panic(err)
	}
	if !dirInfo.IsDir() {
		panic(path.Dir(outputPath) + ": not a directory")
	}

	// download

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// optimize and save

	sizeBefore := float64(len(data))
	data, err = minifySVG(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("optimized: %.1f%% of original\n", (float64(len(data))/sizeBefore)*100)

	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("saved to: " + outputPath)
}
