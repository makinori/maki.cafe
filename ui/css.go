package ui

import (
	"strings"

	"github.com/makinori/maki.cafe/utils"

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

func SCSS(p *Providers, input string) Node {
	var source string

	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		source += line + "\n"
	}

	className := utils.HashString(source)

	p.SharedSCSS[className] = source

	return Class(className)
}

func SCSSEl(p *Providers) Node {
	var source string

	for className, snippet := range p.SharedSCSS {
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
