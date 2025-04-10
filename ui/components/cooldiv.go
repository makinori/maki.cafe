package components

import (
	_ "embed"

	. "github.com/makinori/maki.cafe/common"
	. "github.com/makinori/maki.cafe/ui/render"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CoolDiv(r *RenderContext, children ...Node) Node {
	id := UniqueHashPC()

	r.HeadJS[id] = `const coolDiv = ()=>{
		let e = document.currentScript.parentElement;
		let active = false;

		e.addEventListener("click", ()=>{
			active = !active;
			e.style.background = active ? "blue" : null;
		})
	}`

	return Div(
		Class(SCSS(r, `
			display: flex;
			align-items: center;
			justify-content: center;
			width: 200px;
			height: 200px;
			background: red;
			margin: 16px;
			transition: all 200ms ease-in-out;
			user-select: none;

			&:hover {
				height: 400px;
			}
		`)),
		Text("its working?"),
		Script(Raw(`coolDiv()`)),
	)
}
