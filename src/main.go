package src

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/makinori/maki.cafe/src/config"
	"github.com/makinori/maki.cafe/src/data"
	"github.com/makinori/maki.cafe/src/lint"
	"github.com/makinori/maki.cafe/src/page"
	"github.com/makinori/maki.cafe/src/render"
	"github.com/makinori/maki.cafe/src/util"
	"maragu.dev/gomponents"
)

var (
	//go:embed public
	staticContent embed.FS

	// minifier *minify.M
)

func handlePage(pageFn func(context.Context) gomponents.Group) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := render.InitContext()

		site, err := render.Site(ctx, pageFn(ctx), r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			log.Error("failed to render", "err", err.Error())
		}

		pageBuf := bytes.NewBuffer(nil)

		err = site.Render(pageBuf)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			log.Error("failed to render", "err", err.Error())
			return
		}

		// minSiteBuf := bytes.NewBuffer(nil)
		// err = minifier.Minify("text/html", minSiteBuf, pageBuf)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("failed to minify page"))
		// 	log.Error("failed to minify page", "err", err.Error())
		// 	return
		// }

		go util.HTTPPlausibleEvent(r)

		renderTime := time.Now().Sub(start)

		if util.ENV_IS_DEV {
			log.Debugf("render %s %s", r.URL.Path, renderTime.String())
			lint.LintHTML(pageBuf.Bytes())
		}

		w.Header().Set("X-Render-Time", strings.ReplaceAll(renderTime.String(), "µ", "u"))

		// util.HTTPServeOptimized(w, r, minSiteBuf.Bytes(), ".html")
		util.HTTPServeOptimized(w, r, pageBuf.Bytes(), ".html")
	}
}

func Main() {
	// initialization

	if util.ENV_IS_DEV {
		log.Info("in developer mode")
		log.SetLevel(log.DebugLevel)
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
		http.Redirect(w, r, "mailto:"+config.Email, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("GET /xmpp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "xmpp:"+config.XMPP, http.StatusTemporaryRedirect)
	})

	// register pages

	mux.HandleFunc("GET /{$}", handlePage(page.Index))
	mux.HandleFunc("GET /anime", handlePage(page.Anime))
	mux.HandleFunc("GET /webring", handlePage(page.Webring))

	// register assets

	publicFs, err := fs.Sub(staticContent, "public")
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
	}

	addr := fmt.Sprintf(":%d", port)
	log.Info("listening at " + addr)

	err = http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		log.Fatal("failed to start http server", "err", err.Error())
	}
}
