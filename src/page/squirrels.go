package page

import (
	"context"
	"fmt"

	"github.com/mergestat/timediff"
	"maki.cafe/src/component"
	"maki.cafe/src/config"
	"maki.cafe/src/data"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Squirrels(ctx context.Context) Group {
	columns := 3

	var items []Node

	for _, squirrel := range data.Squirrels.Data {
		title := "squirrel " + timediff.TimeDiff(squirrel.Date)
		items = append(items, component.GridItem(
			title, squirrel.Link,
			fmt.Sprintf("background-image:url(%s)", squirrel.Thumbnail),
		))
	}

	return Group{
		P(
			component.HStack(ctx, []Node{
				Img(Src("/icons/emoji/squirrel.svg"), Height("24")),
				Text("take picture of a squirrel, "),
				I(Text("hopefully everyday")),
			}, "align-items:center"),
		),
		Br(),
		component.Grid(ctx, columns, items, `
			aspect-ratio: 3/4;
		`),
		Br(),
		P(
			Text("find more at my "),
			A(Href(config.SquirrelsURL), Text("mastodon page")),
		),

		// H2(Text("friends")),
		// Br(),
		// Div(
		// 	Class(gridClass),
		// 	webringIcon(ctx, "0x0ade.gif", "0x0a.de"),
		// 	webringIcon(ctx, "lemon.png", "lemon.horse"),
		// 	webringIcon(ctx, "cmtaz.png", "cmtaz.net"),
		// 	webringIcon(ctx, "micaela.gif", "micae.la"),
		// 	webringIcon(ctx, "kayla.gif", "kayla.moe"),
		// 	webringIcon(ctx, "kneesox.png", "kneesox.moe"),
		// 	webringIcon(ctx, "!ironsm4sh.nl", "ironsm4sh.nl"),
		// 	webringIcon(ctx, "skyn3t.gif", "skyn3t.lol"),
		// 	// webringIcon(ctx, "!pony.best", "pony.best"),
		// 	// Div(),
		// ),
		// Br(),
		// Br(),
		// H2(Text("other")),
		// Br(),
		// Div(
		// 	Class(gridClass),
		// 	webringIcon(ctx, "yno.png", "ynoproject.net"),
		// 	webringIcon(ctx, "anonfilly.png", "anonfilly.horse"),
		// ),
		// // https://cyber.dabamos.de/88x31/
		// // https://capstasher.neocities.org/88x31collection
		// // TODO: https://neonaut.neocities.org/cyber/88x31
		// Br(),
		// Br(),
		// // sillyWebring(ctx, gridClass),
		// P(Text("feel free to use my button")),
		// P(Text("although it needs to be updated")),
		// Br(),
		// Div(
		// 	Class(gridClass),
		// 	webringIcon(ctx, "maki.gif", "maki.cafe", "title:or use maki@2x.gif"),
		// ),
		// // Br(),
		// // Br(),
		// // H2(Text("pony")),
		// // Br(),
		// // Div(
		// // 	Class(gridClass),
		// // 	webringIcon(ctx, "!pony.town", "pony.town"),
		// // 	webringIcon(ctx, "!wetmares", "wetmares.org"),
		// // ),
	}
}
