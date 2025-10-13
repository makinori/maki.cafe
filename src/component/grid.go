package component

import (
	"context"
	"fmt"
	"strconv"

	"github.com/makinori/goemo"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Grid(
	ctx context.Context, columns int, items []Node, itemSCSS string,
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
				
				> p {
					font-size: 14px;
				}

				> div {
					width: 100%;
					padding: 0;
					border-radius: 4px;
					background-size: cover;
					`+itemSCSS+`
				}
			}
		`)),
		Group(items),
	)
}

func GridItem(
	name string, href string, itemStyle string, nodes ...Node,
) Node {
	var props []Node

	if name != "" {
		props = append(props, Title(name))
	}
	if href != "" {
		props = append(props, Href(href))
	}

	props = append(props,
		Div(Style(itemStyle)),
		Group(nodes),
	)

	return A(Group(props))
}

func SpriteSheetGrid(
	ctx context.Context, imageURL string, size string, aspectRatio string,
	columns int, items []Node,
) Node {
	return Grid(
		ctx, columns, items,
		`
			aspect-ratio: `+aspectRatio+`;
			background-color: #222;
			background-image: url("`+imageURL+`");
			background-size: `+size+`;
		`,
	)
}

func SpriteSheetGridItem(
	name string, href string, position string, nodes ...Node,
) Node {
	return GridItem(name, href, fmt.Sprintf(
		`background-position: %s;`, position,
	), nodes...)
}
