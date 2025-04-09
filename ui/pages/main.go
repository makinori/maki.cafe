package pages

import (
	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// type Providers struct {
// 	*ui.Providers
// }

func MainPage(p *ui.Providers) Node {
	return Div(
		components.CoolDiv(p),
		components.CoolDiv(p),
	)
}
