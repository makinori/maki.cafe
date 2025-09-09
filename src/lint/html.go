package lint

import (
	"bytes"
	"log/slog"
	"regexp"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
	"golang.org/x/net/html"
)

var cssURLHttpRegexp = regexp.MustCompile(`(?i)(http.+?)["')]`)

func LintHTML(inputHTML string) {
	doc, err := html.Parse(bytes.NewBuffer([]byte(inputHTML)))
	if err != nil {
		slog.Error("failed to parse html", "err", err)
		return
	}

	extHTTPResource := "external http resource"

	parseStyleAttr := func(tag string, style string) {
		l := css.NewLexer(parse.NewInputString(style))
		for {
			tt, text := l.Next()
			switch tt {
			case css.URLToken:
				url := string(text)

				if !strings.Contains(strings.ToLower(url), "http") {
					continue
				}

				matches := cssURLHttpRegexp.FindStringSubmatch(url)
				if len(matches) > 0 {
					url = matches[1]
				}

				slog.Warn(
					extHTTPResource+" in style",
					"tag", tag, "url", url,
				)

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
				parseStyleAttr(tag, attr.Val)
			} else if strings.HasPrefix(
				strings.ToLower(attr.Val), "http",
			) {
				slog.Warn(
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
