package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/disintegration/imaging"
	"maki.cafe/src/spritesheet"
)

// https://www.steamgriddb.com/

const (
	// steamgriddb uses 920x430
	// steam uses 460x215 which is half the size

	// scale        = 0.5
	// bannerWidth  = 920 * scale
	// bannerHeight = 430 * scale

	// however steam capsule images are 231x87
	// they look nicer so lets rescale and crop the steamgriddb images

	bannerWidth   = 231
	bannerHeight  = 87
	bannerPadding = 4

	sheetWidth  = 5
	sheetHeight = 12
)

type game struct {
	// derive info from
	SteamID int
	// override
	Image string
	URL   string
}

type gameCategory struct {
	Title string
	Games []game
}

var games = []gameCategory{
	{
		Title: "metroidbrainia",
		Games: []game{
			{
				SteamID: 210970,
				Image:   "games/the-witness.jpg",
			},
			{SteamID: 553420}, // tunic
			{
				SteamID: 224760,
				Image:   "games/fez.jpg",
			},
			{SteamID: 813230}, // animal well
		},
	},
	{
		Title: "souls",
		Games: []game{
			{SteamID: 570940},  // dark souls
			{SteamID: 367520},  // hollow knight
			{SteamID: 257850},  // hyper light drifer
			{SteamID: 1627720}, // lies of p
		},
	},
	{
		Title: "puzzle",
		Games: []game{
			{
				Image: "games/metroid.png",
				URL:   "https://metroid.nintendo.com",
			},
			{
				SteamID: 620,
				Image:   "games/portal2.png",
			},
			{SteamID: 219890},  // anti chamber
			{SteamID: 375820},  // human resource machine
			{SteamID: 1003590}, // tetris effect
			{SteamID: 427520},  // factorio
			{SteamID: 499180},  // braid anniversary edition
			{SteamID: 504210},  // shenzhen io
			{SteamID: 288160},  // the room
			{
				Image: "games/picross-3d-round-2.jpg",
				URL:   "https://www.youtube.com/watch?v=jA-et0LCpNo",
			},
		},
	},
	{
		Title: "fps",
		Games: []game{
			{SteamID: 220},    // half life 2
			{SteamID: 782330}, // doom eternal
			{SteamID: 976730}, // halo mcc
		},
	},
	{
		Title: "platformer",
		Games: []game{
			{SteamID: 504230}, // celeste
			{SteamID: 17410},  // mirrors edge
			{
				Image: "games/super-mario-odyssey.png",
				URL:   "https://www.nintendo.com/store/products/super-mario-odyssey-switch",
			},
			{
				Image: "games/kirby-and-the-forgotten-land.png",
				URL:   "https://kirbyandtheforgottenland.nintendo.com/",
			},
			{SteamID: 1533420}, // neon white
		},
	},
	{
		Title: "story",
		Games: []game{
			{SteamID: 972660},  // spiritfarer
			{SteamID: 1709170}, // paradise marsh
			{SteamID: 1055540}, // a short hike
			{SteamID: 1332010}, // stray
			// -- new line
			{SteamID: 524220},  // nier automata
			{SteamID: 1113560}, // nier replicant
			{
				Image: "games/earthbound.png",
				URL:   "https://www.youtube.com/watch?v=KXQqhRETBeE",
			},
			{
				Image: "games/mother-3.png",
				URL:   "http://mother3.fobby.net",
			},
			// -- new line
			{SteamID: 303210}, // the beginners guide
			{SteamID: 963000}, // frog detective 1
			{SteamID: 420530}, // one shot
			{SteamID: 319630}, // life is strange
			// -- new line
			{SteamID: 447040},  // watch dogs 2
			{SteamID: 1895880}, // ratchet and clank rift apart
			{SteamID: 253230},  // a hat in time
			{
				Image: "games/catherine-full-body.png",
				URL:   "https://www.catherinethegame.com/fullbody",
			},
			// -- new line
			{
				Image: "games/splatoon-2.png",
				URL:   "https://splatoon.nintendo.com",
			},
		},
	},
	{
		Title: "multiplayer",
		Games: []game{
			{
				SteamID: 2357570,
				Image:   "games/overwatch.png",
			},
			{
				Image: "games/vintage-story.png",
				URL:   "https://www.vintagestory.at",
			},
			{
				Image: "games/minecraft.png",
				URL:   "https://www.betterthanadventure.net",
			},
			{SteamID: 394690}, // tower unite
			{
				Image: "games/fortnite-cropped.png",
				URL:   "https://www.fortnite.com",
			},
			{
				Image: "games/world-of-warcraft.png",
				URL:   "https://worldofwarcraft.blizzard.com/en-us",
			},
			{SteamID: 438100}, // vrchat
		},
	},
	{
		Title: "chill",
		Games: []game{
			{SteamID: 1868140}, // dave the diver
			{
				Image: "games/tropix-2.png",
				// TODO: replace with my link?
				URL: "https://www.tropixgame.com",
			},
			{
				Image: "games/animal-crossing-cropped.png",
				URL:   "https://animalcrossing.nintendo.com",
			},
			{SteamID: 413150}, // stardew valley
			{SteamID: 650700}, // yume nikki
			{
				Image: "games/universal-paperclips.png",
				URL:   "https://www.decisionproblem.com/paperclips/",
			},
		},
	},
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirName := filepath.Dir(filename)

	// generate spritesheet

	var images []image.Image

	for _, category := range games {
		for _, game := range category.Games {
			var imageBytes []byte

			if game.Image != "" {
				var err error
				imageBytes, err = os.ReadFile(filepath.Join(dirName, game.Image))
				must(err)
			} else if game.SteamID == 0 {
				panic("missing image or steam id")
			} else {
				imageURL := fmt.Sprintf(
					// /header.jpg
					"https://cdn.cloudflare.steamstatic.com/steam/apps/%d/capsule_184x69.jpg",
					game.SteamID,
				)
				res, err := http.Get(imageURL)
				must(err)
				defer res.Body.Close()
				imageBytes, err = io.ReadAll(res.Body)
				must(err)
			}

			image, err := imaging.Decode(bytes.NewReader(imageBytes))
			must(err)

			images = append(images, image)
		}
	}

	// write spritesheet

	spritesheetImage, spritesheetCSS, err := spritesheet.Generate(
		bannerWidth, bannerHeight, bannerPadding,
		sheetWidth, sheetHeight, images,
	)
	must(err)

	spritesheetImageBytes := bytes.NewBuffer(nil)
	must(imaging.Encode(spritesheetImageBytes, spritesheetImage, imaging.PNG))

	must(os.WriteFile(
		filepath.Join(dirName, "../../src/public/generated/games.png"),
		spritesheetImageBytes.Bytes(), 0644,
	))

	// write template

	f := jen.NewFile("generated")

	f.PackageComment("generated by cmd/makegames on " +
		time.Now().Format(time.Stamp))

	f.Const().Defs(
		jen.Id("GamesImageURL").Op("=").Lit(
			fmt.Sprintf("/generated/games.png?%d", time.Now().Unix()),
		),
		jen.Id("GamesSize").Op("=").Lit(spritesheetCSS.Size),
		jen.Id("GamesAspectRatio").Op("=").Lit(
			fmt.Sprintf("%d/%d", bannerWidth, bannerHeight),
		),
	)

	f.Type().Id("Game").Struct(
		jen.Id("URL").String(),
		jen.Id("Position").String(),
	)

	f.Type().Id("GameCategory").Struct(
		jen.Id("Title").String(),
		jen.Id("Games").Id("[]Game"),
	)

	i := 0

	renderGame := func(g *jen.Group, game game) {
		var url string

		if game.URL != "" {
			url = game.URL
		} else if game.SteamID == 0 {
			panic("missing steam id or url")
		} else {
			url = fmt.Sprintf(
				"https://store.steampowered.com/app/%d", game.SteamID,
			)
		}

		g.Values(jen.Dict{
			jen.Id("URL"):      jen.Lit(url),
			jen.Id("Position"): jen.Lit(spritesheetCSS.Positions[i]),
		})

		i++
	}

	renderGameCategory := func(g *jen.Group, category gameCategory) {
		g.Values(jen.Dict{
			jen.Id("Title"): jen.Lit(category.Title),
			jen.Id("Games"): jen.Id("[]Game").ValuesFunc(func(g *jen.Group) {
				for _, game := range category.Games {
					renderGame(g, game)
				}
			}),
		})
	}

	f.Var().Defs(
		jen.Id("Games").Op("=").Id("[]GameCategory").ValuesFunc(
			func(g *jen.Group) {
				for _, category := range games {
					renderGameCategory(g, category)
				}
			},
		),
	)

	must(f.Save(filepath.Join(dirName, "../../src/data/generated/games.go")))
}
