package generated

const (
	GamesImageURL    = "{{.ImageURL}}"
	GamesSize        = "{{.Size}}"
	GamesAspectRatio = "{{.AspectRatio}}"
)

type Game struct {
	URL      string
	Position string
}

type GameCategory struct {
	Title string
	Games []Game
}

var (
	Games = []GameCategory{ {{range .Games }}
		{
			Title: "{{.Title}}",
			Games: []Game{ {{range .Games }}
				{URL: "{{.URL}}", Position: "{{.Position}}"},{{end}}
			},
		},{{end}}
	}
)
