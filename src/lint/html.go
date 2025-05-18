package lint

import (
	"bytes"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
	"golang.org/x/net/html"
)

func LintHTML(inputHTML []byte) {
	doc, err := html.Parse(bytes.NewBuffer(inputHTML))
	if err != nil {
		log.Error("failed to parse html", "err", err)
		return
	}

	extHTTPResource := "external http resource"

	parseStyle := func(tag string, style string) {
		l := css.NewLexer(parse.NewInputString(style))
		for {
			tt, text := l.Next()
			switch tt {
			case css.URLToken:
				if strings.Contains(strings.ToLower(string(text)), "http") {
					log.Warn(
						extHTTPResource+" in style",
						"tag", tag, "url", string(text),
					)
				}
			case css.ErrorToken:
				return
			}
		}
	}

	parseAttrs := func(tag string, attrs []html.Attribute) {
		if tag == "a" || len(attrs) == 0 {
			return
		}

		for _, attr := range attrs {
			if strings.ToLower(attr.Key) == "style" {
				parseStyle(tag, attr.Val)
			} else if strings.HasPrefix(
				strings.ToLower(attr.Val), "http",
			) {
				log.Warn(
					extHTTPResource,
					"tag", tag, "url", attr.Val,
				)
			}
		}
	}

	var parseNode func(*html.Node)
	parseNode = func(node *html.Node) {
		for node != nil {
			if node.Type == html.ElementNode {
				parseAttrs(node.Data, node.Attr)
			}
			if node.FirstChild != nil {
				parseNode(node.FirstChild)
			}
			node = node.NextSibling
		}
	}

	parseNode(doc)
}
