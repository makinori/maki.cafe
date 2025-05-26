package render

import (
	"context"
	"log/slog"
	"strings"

	"maki.cafe/src/util"
)

var (
	pageSCSSKey = "pageSCSS"
)

type PageSCSS struct {
	ClassName string
	Snippet   string
}

func initContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, pageSCSSKey, &[]PageSCSS{})
	return ctx
}

// returns class name and injects scss into page
func SCSS(ctx context.Context, snippet string) string {
	pageSCSS, ok := ctx.Value(pageSCSSKey).(*[]PageSCSS)
	if !ok {
		slog.Error("failed to get page scss from context")
		return ""
	}

	// TODO: doesnt consider whitespace
	className := util.HashString(snippet)

	for _, scss := range *pageSCSS {
		if scss.ClassName == className {
			return className
		}
	}

	*pageSCSS = append(*pageSCSS, PageSCSS{
		ClassName: className,
		Snippet:   snippet,
	})

	return className
}

func getPageSCSS(ctx context.Context) string {
	pageSCSS, ok := ctx.Value(pageSCSSKey).(*[]PageSCSS)
	if !ok {
		slog.Error("failed to get page scss from context")
		return ""
	}

	var source string

	for _, scss := range *pageSCSS {
		source += "." + scss.ClassName + "{" + scss.Snippet + "}"
	}

	source = strings.TrimSpace(source)

	return source
}
