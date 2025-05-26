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
	"strings"
	"text/template"
	"time"

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

type gameInput struct {
	// derive info from
	SteamID int
	// override
	Image string
	URL   string
}

type gameCategory struct {
	Title string
	Games []gameInput
}

// TODO: add the room, shenzhen io, add favorites section, make wider?

var games = []gameCategory{
	{
		Title: "metroidbrania",
		Games: []gameInput{
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
		Games: []gameInput{
			{SteamID: 1627720}, // lies of p
			{SteamID: 570940},  // dark souls
			{SteamID: 257850},  // hyper light drifer
			{SteamID: 367520},  // hollow knight
		},
	},
	{
		Title: "puzzle",
		Games: []gameInput{
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
			{
				Image: "games/picross-3d-round-2.jpg",
				URL:   "https://www.youtube.com/watch?v=jA-et0LCpNo",
			},
		},
	},
	{
		Title: "fps",
		Games: []gameInput{
			{SteamID: 220},    // half life 2
			{SteamID: 782330}, // doom eternal
			{SteamID: 976730}, // halo mcc
		},
	},
	{
		Title: "platformer",
		Games: []gameInput{
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
		Games: []gameInput{
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
		Games: []gameInput{
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
		Games: []gameInput{
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

type gameOutput struct {
	URL      string
	Position string
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

	// prepare template data

	var allOutputGames []map[string]any

	i := 0

	for _, category := range games {
		var outputGames []gameOutput

		for _, game := range category.Games {
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

			outputGames = append(outputGames,
				gameOutput{
					URL: url,
					Position: strings.ReplaceAll(
						spritesheetCSS.Positions[i], "%", "%",
					),
				},
			)
			i++
		}

		allOutputGames = append(allOutputGames, map[string]any{
			"Title": category.Title,
			"Games": outputGames,
		})
	}

	// write template

	templateBytes, err := os.ReadFile(filepath.Join(dirName, "_template.go"))
	must(err)

	t, err := template.New("games").Parse(string(templateBytes))
	must(err)

	templateData := map[string]any{
		"ImageURL":    "/generated/games.png",
		"Size":        spritesheetCSS.Size, // css
		"AspectRatio": fmt.Sprintf("%d/%d", bannerWidth, bannerHeight),
		"Games":       allOutputGames,
	}

	templateOut := bytes.NewBuffer(nil)
	must(t.Execute(templateOut, templateData))

	must(os.WriteFile(
		filepath.Join(dirName, "../../src/data/generated/games.go"),
		[]byte(
			"// generated by cmd/makegames on "+
				time.Now().Format(time.Stamp)+"\n\n"+
				templateOut.String(),
		), 0644,
	))
}
