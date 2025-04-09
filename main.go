package main

import (
	"fmt"

	"github.com/makinori/maki.cafe/ui"
	"github.com/makinori/maki.cafe/ui/pages"

	. "maragu.dev/gomponents"
)

func main() {
	p := ui.NewProviders()

	html := Group{
		ui.Layout(&p, pages.MainPage(&p)),
	}.String()

	fmt.Println(html)
}
