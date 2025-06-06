package src

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"maki.cafe/src/config"
	"maki.cafe/src/data"
	"maki.cafe/src/lint"
	"maki.cafe/src/page"
	"maki.cafe/src/render"
	"maki.cafe/src/template"
	"maki.cafe/src/util"
	"maragu.dev/gomponents"
)

var (
	//go:embed public
	staticContent embed.FS
	//go:embed 1x1.gif
	gif1x1 []byte
)

func handlePage(pageFn func(context.Context) gomponents.Group) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		html, err := template.RenderPage(pageFn, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to render"))
			slog.Error("failed to render", "err", err.Error())
		}

		renderTime := time.Now().Sub(start)

		if util.ENV_IS_DEV {
			slog.Debug("render", "path", r.URL.Path, "time", renderTime)
			lint.LintHTML(html)
		}

		w.Header().Set("X-Render-Time", strings.ReplaceAll(renderTime.String(), "µ", "u"))

		util.HTTPServeOptimized(w, r, html, ".html")
	}
}

func handleNotabotGif(w http.ResponseWriter, r *http.Request) {
	// respond immediately
	w.Header().Add("Cache-Control", "no-store")
	w.Write(gif1x1)

	go func() {
		data.AddOneToCounter(r)
		util.HTTPPlausibleEventFromImg(r)
	}()
}

func Main() {
	// initialization

	if util.ENV_IS_DEV {
		slog.Info("in developer mode")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	data.InitData()
	render.InitSass()

	// no need to cause sass and gomponents are already mostly minified
	// render.InitMinifier()

	// register api

	mux := http.NewServeMux()

	mux.HandleFunc("GET /email", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "mailto:"+config.Email, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("GET /xmpp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "xmpp:"+config.XMPP, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("GET /notabot.gif", handleNotabotGif)

	// register pages

	hP := handlePage
	mux.HandleFunc("GET /{$}", hP(page.Index))
	mux.HandleFunc("GET /webring", hP(page.Webring))
	mux.HandleFunc("GET /fav/anime", hP(page.FavAnime))
	mux.HandleFunc("GET /fav/games", hP(page.FavGames))

	// register assets

	mux.HandleFunc(
		"GET /cache/{file...}", util.HTTPFileServerOptimized(
			os.DirFS("cache/public"),
		),
	)

	publicFs, err := fs.Sub(staticContent, "public")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

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
			slog.Error(err.Error())
			os.Exit(1)
		}
	}

	addr := fmt.Sprintf(":%d", port)
	slog.Info("listening at " + addr)

	err = http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		slog.Error("failed to start http server", "err", err.Error())
		os.Exit(1)
	}
}
