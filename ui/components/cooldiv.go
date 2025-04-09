package components

import (
	"github.com/makinori/maki.cafe/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CoolDiv(r *ui.Renderer, children ...Node) Node {
	return Div(
		ui.SCSS(r, `
			$what: 200px;

			display: flex;
			align-items: center;
			justify-content: center;

			:hover {
				height: $what;
			}
		`),
		Text("its working?"),
	)
}
