package src

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/makinori/maki.cafe/src/common"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/page"
	"github.com/makinori/maki.cafe/src/render"
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
		start := time.Now()

		template, err := template.Site(pageFn(), r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			log.Println("failed to render: " + err.Error())
		}

		pageBuf := bytes.NewBuffer(nil)

		err = template.Render(pageBuf)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			log.Println("failed to render: " + err.Error())
			return
		}

		// minSiteBuf := bytes.NewBuffer(nil)
		// err = minifier.Minify("text/html", minSiteBuf, pageBuf)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("failed to minify page"))
		// 	log.Println("failed to minify page: " + err.Error())
		// 	return
		// }

		go util.HTTPPlausibleEvent(r)

		renderTime := time.Now().Sub(start)

		if util.ENV_IS_DEV {
			log.Println("render " + r.URL.Path + " " + renderTime.String())
		}

		w.Header().Set("X-Render-Time", strings.ReplaceAll(renderTime.String(), "µ", "u"))

		// util.HTTPServeOptimized(w, r, minSiteBuf.Bytes(), ".html")
		util.HTTPServeOptimized(w, r, pageBuf.Bytes(), ".html")
	}
}

func Main() {
	// initialization

	if util.ENV_IS_DEV {
		log.Println("in developer mode")
	}

	data.InitData()
	render.InitSass()

	// no need to cause sass and gomponents are already mostly minified
	// minifier = minify.New()
	// minifier.Add("text/css", &css.Minifier{})
	// minifier.Add("text/html", &html.Minifier{
	// 	KeepDocumentTags:    true,
	// 	KeepQuotes:          true,
	// 	KeepDefaultAttrVals: true,
	// 	// TODO: minifier removes character entities
	// })

	// register api

	mux := http.NewServeMux()

	mux.HandleFunc("GET /email", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "mailto:"+common.Email, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("GET /xmpp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "xmpp:"+common.XMPP, http.StatusTemporaryRedirect)
	})

	// register pages

	mux.HandleFunc("GET /{$}", handlePage(page.Index))
	mux.HandleFunc("GET /anime", handlePage(page.Anime))

	// register assets

	publicFs, err := fs.Sub(staticContent, "public")
	if err != nil {
		log.Fatalln(err)
	}

	// http.FileServerFS(publicFs)
	mux.HandleFunc(
		"GET /{file...}", util.HTTPFileServerOptimized(publicFs),
	)

	// middleware

	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", runtime.Version()) // hell yeah
		mux.ServeHTTP(w, r)
	})

	// listen

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

	err = http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		log.Fatalln("failed to start http server: " + err.Error())
	}
}
