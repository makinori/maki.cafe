package components

import (
	"github.com/makinori/maki.cafe/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CoolDiv(p *ui.Providers, children ...Node) Node {
	return Div(
		ui.CSS(p, `
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
