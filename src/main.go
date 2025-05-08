package src

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/makinori/maki.cafe/src/page"
	"github.com/makinori/maki.cafe/src/template"
	"github.com/makinori/maki.cafe/src/util"
	"maragu.dev/gomponents"
)

var (
	//go:embed public
	staticContent embed.FS

	// minifier *minify.M
)

func handlePage(pageFn func() gomponents.Group) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// render

		pageBuf := bytes.NewBuffer(nil)

		err := template.Site(pageFn()).Render(pageBuf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			log.Println("failed to render: " + err.Error())
			return
		}

		// minify and write

		// minSiteBuf := bytes.NewBuffer(nil)

		// err = minifier.Minify("text/html", minSiteBuf, pageBuf)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("failed to minify page"))
		// 	log.Println("failed to minify page: " + err.Error())
		// 	return
		// }

		go util.HTTPPlausibleEvent(r)

		// util.HTTPServeOptimized(w, r, minSiteBuf.Bytes())
		util.HTTPServeOptimized(w, r, pageBuf.Bytes(), ".html")
	}
}

func Main() {
	// register minifiers

	// minifier = minify.New()
	// minifier.Add("text/css", &css.Minifier{})
	// minifier.Add("text/html", &html.Minifier{
	// 	KeepDocumentTags:    true,
	// 	KeepQuotes:          true,
	// 	KeepDefaultAttrVals: true,
	// 	// TODO: minifier removes character entities
	// })

	// register page handles

	http.HandleFunc("GET /{$}", handlePage(page.Index))

	// register assets

	publicFs, err := fs.Sub(staticContent, "public")
	if err != nil {
		log.Fatalln(err)
	}

	// http.FileServerFS(publicFs)
	http.HandleFunc(
		"GET /{file...}", util.HTTPFileServerOptimized(publicFs),
	)

	port := 8080

	portStr := os.Getenv("PORT")
	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			log.Fatalln(err)
		}
	}

	addr := fmt.Sprintf(":%d", port)
	log.Println("listening at " + addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln("failed to start http server: " + err.Error())
	}
}
