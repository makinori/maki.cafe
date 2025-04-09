package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// prop drilling
// TODO: can we improve this?

type Providers struct {
	SharedCSS map[string]string
}

func NewProviders() Providers {
	ensureSass()

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
