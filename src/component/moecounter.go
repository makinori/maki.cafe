package component

import (
	"context"
	"fmt"

	"github.com/makinori/goemo"
	"maki.cafe/src/data"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MoeCounter(ctx context.Context) Node {
	var nCharNodes Group

	length := 6
	chars := fmt.Sprintf("%d", data.ReadCounter())

	anchorClass := goemo.SCSS(ctx, `
		margin: 0;
		padding: 0;
		background: none;
		gap: 0;
		
		> img {
			// original-new 150px
			// rule34 100px
			height: 75px;
			image-rendering: pixelated;
			&.padding {
				opacity: 0.5;
			}
		}
			
	`)

	for range max(0, length-len(chars)) {
		nCharNodes = append(nCharNodes, Img(
			Class("padding"),
			Src("/moecounter/rule34/0.gif"),
		))
	}
	for _, char := range chars {
		nCharNodes = append(nCharNodes, Img(
			Src(fmt.Sprintf("/moecounter/rule34/%c.gif", char)),
		))
	}

	return A(
		Class(anchorClass),
		Href("https://github.com/journey-ad/Moe-Counter"),
		Title("since june 2025"), // specifically june 1st
		nCharNodes,
	)
}
