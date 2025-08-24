package component

import (
	"context"
	"strings"

	"github.com/makinori/emgotion"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func HStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return Div(
		Classes{
			emgotion.SCSS(ctx, `
				display: flex;
				flex-direction: row;
				gap: 8px;
			`): true,
			emgotion.SCSS(ctx, strings.Join(extraSCSS, "\n")): true,
		},
		Group(args),
	)
}

func VStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return Div(
		Classes{
			emgotion.SCSS(ctx, `
				display: flex;
				flex-direction: column;
				gap: 8px;
			`): true,
			emgotion.SCSS(ctx, strings.Join(extraSCSS, "\n")): true,
		},
		Group(args),
	)
}
