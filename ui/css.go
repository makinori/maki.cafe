package ui

import (
	"strings"

	"github.com/makinori/maki.cafe/utils"

	sass "github.com/bep/godartsass/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CSS(p *Providers, input string) Node {
	var source string

	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		source += line + "\n"
	}

	className := utils.HashString(source)

	p.SharedCSS[className] = source

	return Class(className)
}

func CSSEl(p *Providers) Node {
	var source string

	for className, snippet := range p.SharedCSS {
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
