package render

import (
	"context"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/makinori/maki.cafe/src/util"
)

var (
	pageSCSSKey = "pageSCSS"
)

func InitContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, pageSCSSKey, map[string]string{})
	return ctx
}

// returns class name and injects scss into page
func SCSS(ctx context.Context, scss string) string {
	pageSCSS, ok := ctx.Value(pageSCSSKey).(map[string]string)
	if !ok {
		log.Error("failed to get page scss from context")
		return ""
	}

	className := util.HashString(scss)

	pageSCSS[className] = scss

	return className
}

func getPageSCSS(ctx context.Context) string {
	pageSCSS, ok := ctx.Value(pageSCSSKey).(map[string]string)
	if !ok {
		log.Error("failed to get page scss from context")
		return ""
	}

	var source string

	for className, snippet := range pageSCSS {
		source += "." + className + "{" + snippet + "}"
	}

	source = strings.TrimSpace(source)

	return source
}
