package component

import (
	"context"
	"fmt"
	"strconv"

	"github.com/makinori/maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SpriteSheetGrid(
	ctx context.Context, imageURL string, size string, aspectRatio string,
	columns int, items []Node,
) Node {
	return Div(
		Class(render.SCSS(ctx, `
			display: grid;
			grid-template-columns: repeat(`+strconv.Itoa(columns)+`, 1fr);
			grid-gap: 8px;
			
			> a {
				background-image: url("`+imageURL+`");
				background-size: `+size+`;
				aspect-ratio: `+aspectRatio+`;
				padding: 0;

				display: flex;
				flex-direction: column;
				gap: 4px;
				background-color: #222;
				
				> p {
					font-size: 18px;
				}
			}
		`)),
		Group(items),
	)
}

func SpriteSheetGridItem(
	ctx context.Context, name string, href string,
	position string, nodes ...Node,
) Node {
	var props []Node

	if name != "" {
		props = append(props, Title(name))
	}
	if href != "" {
		props = append(props, Href(href))
	}

	props = append(props,
		Style(fmt.Sprintf(
			`background-position: %s;`, position,
		)),
		Group(nodes),
	)

	return A(Group(props))
}
