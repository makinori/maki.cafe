package pages

import (
	. "github.com/makinori/maki.cafe/ui/components"
	. "github.com/makinori/maki.cafe/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MainPage(r *RenderContext) Node {
	return Div(
		Class(SCSS(r, `
			display: flex;
			flex-direction: row;
		`)),
		CoolDiv(r),
		CoolDiv(r, Attr("test2"), Class("test")),
	)
}
