package components

import (
	_ "embed"

	. "github.com/makinori/maki.cafe/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed cooldiv.scss
var styles string

func CoolDiv(r *Renderer, children ...Node) Node {
	return Div(
		Classes([]string{SCSS(r, styles), "box"}),
		Text("its working?"),
	)
}
