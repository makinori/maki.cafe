package main

import (
	"os"

	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/pages"
)

func main() {
	html := ui.Render(pages.MainPage)
	os.WriteFile("output.html", []byte(html), 0644)
}
