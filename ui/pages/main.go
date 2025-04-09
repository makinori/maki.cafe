package pages

import (
	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MainPage(r *ui.Renderer) Node {
	return Div(
		components.CoolDiv(r),
		components.CoolDiv(r),
	)
}
