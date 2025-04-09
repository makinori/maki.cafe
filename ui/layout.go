package ui

import (
	sass "github.com/bep/godartsass/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// prop drilling
// TODO: can we improve this?

type Providers struct {
	SharedCSS map[string]string
}

var sassTranspiler *sass.Transpiler

func NewProviders() Providers {
	if sassTranspiler == nil {
		var err error
		sassTranspiler, err = sass.Start(sass.Options{})
		if err != nil {
			panic(err)
		}
	}

	return Providers{
		SharedCSS: map[string]string{},
	}
}

func Layout(p *Providers, children ...Node) Node {
	return HTML(
		Head(
			TitleEl(Text("maki")),
			CSSEl(p),
		),
		Body(children...),
	)
}
