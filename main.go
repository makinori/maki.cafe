package main

import (
	"os"

	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/pages"
	"github.com/makinori/maki.cafe/ui/render"
)

func main() {
	html := render.Render(ui.Layout, pages.MainPage)
	os.WriteFile("output.html", []byte(html), 0644)
}
