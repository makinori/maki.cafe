package page

import (
	"context"

	. "maragu.dev/gomponents"
	// . "maragu.dev/gomponents/html"
)

func DlBlender(ctx context.Context) Group {
	return dlPage(
		ctx, "big/blender", "/blender",
		"blender goodies to download",
		"things i've made or ported and such",
	)
}
