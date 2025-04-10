package components

import (
	_ "embed"

	. "github.com/makinori/maki.cafe/common"
	. "github.com/makinori/maki.cafe/ui"
	. "github.com/makinori/maki.cafe/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed cooldiv.scss
var styles string

func CoolDiv(r *RenderContext, children ...Node) Node {
	id := UniqueHashPC()

	r.JS[id] = `
		for (const el of document.querySelectorAll("#` + id + `")) {
			el.addEventListener("click", ()=>{
				alert(el);
			});
		} 
	`

	return Div(
		ID(id),
		Classes([]string{SCSS(r, styles), "box"}),
		Text("its working?"),
	)
}
