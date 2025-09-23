package component

import (
	"context"
	"fmt"
	"strconv"

	"github.com/makinori/goemo"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SpriteSheetGrid(
	ctx context.Context, imageURL string, size string, aspectRatio string,
	columns int, items []Node,
) Node {
	return Div(
		Class(goemo.SCSS(ctx, `
			display: grid;
			grid-template-columns: repeat(`+strconv.Itoa(columns)+`, 1fr);
			grid-gap: 8px;
			
			> a {
				padding: 0;
				display: flex;
				flex-direction: column;
				gap: 4px;
				background-color: transparent;

				> div {
					width: 100%;
					padding: 0;
					aspect-ratio: `+aspectRatio+`;
					background-color: #222;
					background-image: url("`+imageURL+`");
					background-size: `+size+`;
					border-radius: 4px;
				}
				
				> p {
					font-size: 14px;
				}
			}
		`)),
		Group(items),
	)
}

func SpriteSheetGridItem(
	name string, href string, position string, nodes ...Node,
) Node {
	var props []Node

	if name != "" {
		props = append(props, Title(name))
	}
	if href != "" {
		props = append(props, Href(href))
	}

	props = append(props,
		Div(Style(fmt.Sprintf(
			`background-position: %s;`, position,
		))),
		Group(nodes),
	)

	return A(Group(props))
}
