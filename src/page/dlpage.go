package page

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
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
)

func dlPage(
	ctx context.Context,
	markdownDir, dlPath, title, description, icon string,
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

	for _, mdFile := range allFiles {
		ext := filepath.Ext(mdFile.Name())
		if ext != ".md" {
			continue
		}

		id := strings.TrimSuffix(mdFile.Name(), ext)

		htmlError := func(err error) {
			if err != nil {
				slog.Error(err.Error())
				nodes = append(nodes,
					P(Text("failed to render: "+mdFile.Name())), Br(),
				)
			}
		}

		markdown, err := os.ReadFile(
			filepath.Join(markdownDir, mdFile.Name()),
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
			P(Raw(html)),
			Br(),
		)

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
			nodes = append(nodes,
				Img(Src(dlPath+"/"+image), Width("100%")),
				Br(),
				Br(),
				Div(Class(spacer)),
			)
		}

	}

	return Group{
		component.IconHeader(ctx, title, icon),
		Br(),
		P(Text(description)),
		Br(),
		nodes,
	}
}
