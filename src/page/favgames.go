package page

import (
	"context"

	"maki.cafe/src/component"
	"maki.cafe/src/data/generated"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func FavGames(ctx context.Context) Group {
	var nodes Group

	for _, category := range generated.Games {
		nodes = append(nodes, H1(Text(category.Title)), Br())

		var items []Node

		for _, game := range category.Games {
			// fmt.Println(game)
			items = append(items, component.SpriteSheetGridItem(
				"", game.URL, game.Position,
			))
		}

		nodes = append(nodes,
			component.SpriteSheetGrid(ctx,
				generated.GamesImageURL, generated.GamesSize,
				generated.GamesAspectRatio, 4, items,
			), Br(),
		)
	}

	// remove last br
	nodes = nodes[:len(nodes)-1]

	return Group{
		nodes,
	}
}
