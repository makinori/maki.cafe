package component

import (
	"context"
	"fmt"
	"strconv"

	"github.com/makinori/maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func SpriteSheetGridItem(
	ctx context.Context, name string, href string,
	imageURL string, size string, position string,
	aspectRatio string, nodes ...Node,
) Node {
	var props []Node

	if name != "" {
		props = append(props, Title(name))
	}

	var extraClass string
	var itemExtraCSS string
	if aspectRatio != "" {
		itemExtraCSS += "aspect-ratio: " + aspectRatio + ";"
	}
	if itemExtraCSS != "" {
		extraClass = render.SCSS(ctx, `
			> div {
				`+itemExtraCSS+`
			}
		`)
	}

	props = append(props,
		Href(href),
		Classes{render.SCSS(ctx, `
			padding: 0;
			display: flex;
			flex-direction: column;
			gap: 4px;
			background: transparent;

			> div {
				width: 100%;
				background-color: #222;
			}

			> p {
				font-size: 18px;
			}
		`): true, extraClass: true},
		Div(
			Class(render.SCSS(ctx, `
				background-image: url("`+imageURL+`");
				background-size: `+size+`;
			`)),
			Style(fmt.Sprintf(
				`background-position: %s;`, position,
			)),
		),
		Group(nodes),
	)

	return A(props...)
}

func SpriteSheetGrid(
	ctx context.Context, columns int, items []Node,
) Node {
	return Div(
		Class(render.SCSS(ctx, `
			display: grid;
			grid-template-columns: repeat(`+strconv.Itoa(columns)+`, 1fr);
			grid-gap: 8px;
		`)),
		Group(items),
	)
}
