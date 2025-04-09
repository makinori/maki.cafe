package main

import (
	"fmt"

	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/pages"
)

func main() {
	html := ui.Render(pages.MainPage)
	fmt.Println(html)
}
