package main

import (
	"os"

	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/pages"
	"github.com/makinori/maki.cafe/ui/render"
)

func main() {
	html := render.Render(ui.Layout, pages.MainPage)

	os.RemoveAll("build")
	os.MkdirAll("build", 0755)

	os.CopyFS("build", os.DirFS("public"))

	os.WriteFile("build/index.html", []byte(html), 0644)
}
