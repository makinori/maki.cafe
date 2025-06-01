package component

import (
	"context"
	"fmt"
	"strconv"

	"maki.cafe/src/data"
	"maki.cafe/src/render"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MoeCounter(ctx context.Context) Node {
	var nCharNodes Group
	nChars := strconv.FormatUint(data.ReadCounter(), 10)

	aClass := render.SCSS(ctx, `
		margin: 0;
		padding: 0;
		background: none;
		> img {
			// height: 150px; // default height
			height: 75px;
			image-rendering: pixelated;
		}
	`)

	for _, char := range nChars {
		nCharNodes = append(nCharNodes, Img(
			Src(fmt.Sprintf("/moecounter/%c.gif", char)),
		))
	}

	return A(
		Class(aClass),
		Href("https://github.com/journey-ad/Moe-Counter"),
		nCharNodes,
	)
}
