package component

import (
	"context"

	"github.com/makinori/foxlib/foxhtml"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IconHeader(
	ctx context.Context, header Node, icon string,
) Node {
	return foxhtml.HStack(ctx,
		foxhtml.StackSCSS(`
			gap: 16px;
			align-items: center;
		`),
		Img(Src(icon), Height("32")),
		header,
	)
}
