package render

import (
	"errors"
	"os"
	"strings"

	sass "github.com/bep/godartsass/v2"
	"golang.org/x/exp/slog"
)

var (
	sassTranspiler *sass.Transpiler
)

type SassImport struct {
	Filename string
	Content  string
}

type embeddedImportResolver struct {
	imports []SassImport
}

func (importResolver embeddedImportResolver) CanonicalizeURL(url string) (string, error) {
	return "embed://" + url, nil
}

func (importResolver embeddedImportResolver) Load(canonicalizedURL string) (sass.Import, error) {
	if !strings.HasPrefix(canonicalizedURL, "embed://") {
		return sass.Import{}, errors.New("invalid url")
	}

	filename := canonicalizedURL[8:]

	for _, sassImport := range importResolver.imports {
		if sassImport.Filename == filename {
			sourceSyntax := sass.SourceSyntaxSCSS
			switch {
			case strings.HasPrefix(sassImport.Filename, ".sass"):
				sourceSyntax = sass.SourceSyntaxSASS
			case strings.HasPrefix(sassImport.Filename, ".css"):
				sourceSyntax = sass.SourceSyntaxCSS
			}

			return sass.Import{
				SourceSyntax: sourceSyntax,
				Content:      sassImport.Content,
			}, nil
		}
	}

	return sass.Import{}, errors.New("failed to find " + filename)
}

func RenderSass(source string, imports ...SassImport) (string, error) {
	res, err := sassTranspiler.Execute(sass.Args{
		ImportResolver: embeddedImportResolver{
			imports: imports,
		},
		Source:          source,
		OutputStyle:     sass.OutputStyleCompressed,
		SourceSyntax:    sass.SourceSyntaxSCSS,
		EnableSourceMap: false,
	})

	if err != nil {
		return "", errors.New("failed to compile scss")
	}

	return res.CSS, nil
}

func InitSass() {
	var err error
	sassTranspiler, err = sass.Start(sass.Options{})
	if err != nil {
		slog.Error("failed to start sass transpiler", "err", err.Error())
		os.Exit(1)
	}
}
