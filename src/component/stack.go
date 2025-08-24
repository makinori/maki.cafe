package component

import (
	"context"
	"strings"

	"github.com/makinori/emgotion"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func stack(
	ctx context.Context, flexDir string, args []Node, extraSCSS ...string,
) Node {
	class := emgotion.SCSS(ctx, `
		display: flex;
		flex-direction: `+flexDir+`;
		gap: 8px;
	`)

	if len(extraSCSS) > 0 {
		class += " " + emgotion.SCSS(ctx, strings.Join(extraSCSS, "\n"))
	}

	return Div(
		Class(class),
		Group(args),
	)
}
func HStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return stack(ctx, "row", args, extraSCSS...)
}

func VStack(ctx context.Context, args []Node, extraSCSS ...string) Node {
	return stack(ctx, "column", args, extraSCSS...)
}
