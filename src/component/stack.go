package component

import (
	"context"
	"strings"

	"maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func HStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return Div(
		Classes{
			render.SCSS(ctx, `
				display: flex;
				flex-direction: row;
				gap: 8px;
			`): true,
			render.SCSS(ctx, strings.Join(extraSCSS, "\n")): true,
		},
		Group(args),
	)
}

func VStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return Div(
		Classes{
			render.SCSS(ctx, `
				display: flex;
				flex-direction: column;
				gap: 8px;
			`): true,
			render.SCSS(ctx, strings.Join(extraSCSS, "\n")): true,
		},
		Group(args),
	)
}
