package components

import (
	"github.com/makinori/maki.cafe/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CoolDiv(r *ui.Renderer, children ...Node) Node {
	return Div(
		ui.SCSS(r, `
			display: flex;
			align-items: center;
			justify-content: center;
			width: 200px;
			height: 200px;
			background: red;
			margin: 16px;
			transition: all 200ms ease-in-out;

			&:hover {
				height: 400px;
			}
		`),
		Text("its working?"),
	)
}
