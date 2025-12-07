package component

import (
	"context"

	"github.com/makinori/foxlib/foxhtml"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func IconHeader(ctx context.Context, title, icon string) Node {
	heading := H1(Text(title))
	if icon == "" {
		return heading
	}

	return foxhtml.HStack(ctx,
		foxhtml.StackSCSS(`
			gap: 16px;
			align-items: center;
		`),
		Img(Src(icon), Height("32")),
		heading,
	)
}
