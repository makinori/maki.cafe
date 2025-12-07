package page

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/makinori/foxlib/foxcss"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"maki.cafe/src/component"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var (
	dlPageMd = goldmark.New(
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)

	dlPageIDPrefixRegexp = regexp.MustCompile(`^[0-9]+-`)
)

func dlPage(
	ctx context.Context,
	markdownDir, dlPath, icon,
	title, description string,
	descending bool,
) Group {
	allFiles, err := os.ReadDir(markdownDir)
	if err != nil {
		return Group{
			H2(Text("failed to get markdown files")),
			H2(Text("try again later :(")),
		}
	}

	var nodes Group

	spacer := foxcss.Class(ctx, `
		margin: 32px 0;
	`)

	var mdFiles []string
	for _, file := range allFiles {
		ext := filepath.Ext(file.Name())
		if ext != ".md" {
			continue
		}
		mdFiles = append(mdFiles, strings.TrimSuffix(file.Name(), ext))
	}

	sort.Strings(mdFiles)
	if descending {
		slices.Reverse(mdFiles)
	}

	for _, id := range mdFiles {
		fileName := id + ".md"
		id = dlPageIDPrefixRegexp.ReplaceAllString(id, "")

		htmlError := func(err error) {
			if err != nil {
				slog.Error(err.Error())
				nodes = append(nodes,
					P(Text("failed to render: "+id)), Br(),
				)
			}
		}

		markdown, err := os.ReadFile(
			filepath.Join(markdownDir, fileName),
		)
		if err != nil {
			htmlError(err)
			continue
		}

		doc := dlPageMd.Parser().Parse(text.NewReader(markdown))
		if doc == nil {
			htmlError(errors.New("failed to parse markdown"))
			continue
		}

		meta := doc.OwnerDocument().Meta()

		htmlBuf := bytes.NewBuffer(nil)
		err = dlPageMd.Renderer().Render(htmlBuf, markdown, doc)
		if err != nil {
			htmlError(err)
			continue
		}

		html := htmlBuf.String()
		html = strings.ReplaceAll(html, "<a", `<a class="muted"`)

		nodes = append(nodes,
			Div(Class(spacer)),
			H2(ID(id), A(
				Href("#"+id),
				Class("plain"),
				Text(meta["name"].(string)),
			)),
			P(B(Text("last updated: "+meta["updated"].(string)))),
			Br(),
		)

		if html != "" {
			nodes = append(nodes,
				Raw(html),
				Br(),
			)
		}

		files, ok := meta["files"].([]any)
		if ok {
			for _, file := range files {
				fileName := file.(string)
				nodes = append(nodes,
					A(Text(fileName), Href(dlPath+"/"+fileName)),
					Text(" "),
				)
			}
			nodes = append(nodes, Br(), Br())
		}

		image, ok := meta["image"].(string)
		if ok {
			imgSrc := dlPath + "/" + image
			nodes = append(nodes,
				A(
					Class("plain"),
					Href(imgSrc),
					Img(
						Src(imgSrc), Width("100%"),
						Style("border-radius:8px"),
					),
				),
				Br(),
				Br(),
				Div(Class(spacer)),
			)
		}

	}

	var header Node
	if icon == "" {
		header = H1(Text(title))
	} else {
		header = component.IconHeader(ctx, H1(Text(title)), icon)
	}

	return Group{
		header,
		Br(),
		P(Text(description)),
		Br(),
		nodes,
	}
}
