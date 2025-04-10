package pages

import (
	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MainPage(r *ui.RenderContext) Node {
	return Div(
		Class(ui.SCSS(r, `
			display: flex;
			flex-direction: row;
		`)),
		components.CoolDiv(r),
		components.CoolDiv(r),
	)
}
