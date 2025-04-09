package ui

import (
	"strings"

	"github.com/makinori/maki.cafe/shared"

	sass "github.com/bep/godartsass/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var sassTranspiler *sass.Transpiler

func ensureSass() {
	if sassTranspiler != nil {
		return
	}

	var err error
	sassTranspiler, err = sass.Start(sass.Options{})

	if err != nil {
		panic(err)
	}
}

func SCSS(r *Renderer, input string) string {
	var source string

	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		source += line + "\n"
	}

	className := shared.HashString(source)

	r.SharedSCSS[className] = source

	return className
}

func SCSSEl(r *Renderer, extraScss ...string) Node {
	var source string

	for _, snippet := range extraScss {
		source += snippet + "\n"
	}

	for className, snippet := range r.SharedSCSS {
		source += "." + className + "{" + snippet + "} "
	}

	source = strings.TrimSpace(source)

	res, err := sassTranspiler.Execute(sass.Args{
		Source:          source,
		OutputStyle:     sass.OutputStyleCompressed,
		SourceSyntax:    sass.SourceSyntaxSCSS,
		EnableSourceMap: false,
	})

	if err != nil {
		panic(err)
	}

	return StyleEl(Raw(res.CSS))
}
