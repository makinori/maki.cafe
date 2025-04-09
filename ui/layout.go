package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// prop drilling
// TODO: can we improve this?

type Providers struct {
	SharedSCSS map[string]string
}

func NewProviders() Providers {
	ensureSass()

	return Providers{
		SharedSCSS: map[string]string{},
	}
}

func Layout(p *Providers, children ...Node) Node {
	return HTML(
		Head(
			TitleEl(Text("maki")),
			SCSSEl(p),
		),
		Body(children...),
	)
}
