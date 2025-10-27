package src

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bep/godartsass/v2"
	"github.com/makinori/goemo"
	"github.com/makinori/goemo/emohttp"
	"maki.cafe/src/config"
	"maki.cafe/src/data"
	"maki.cafe/src/lint"
	"maki.cafe/src/page"
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

		renderTime := time.Since(start)
		renderTimeStr := util.ShortDuration(renderTime)

		if util.ENV_IS_DEV {
			slog.Debug("render", "path", r.URL.Path, "time", renderTimeStr)
			lint.LintHTML(html)
		}

		// w.Header().Set("X-Render-Time", strings.ReplaceAll(renderTimeStr, "µ", "u"))
		html = strings.ReplaceAll(html, "{{.RenderTime}}", renderTimeStr)

		emohttp.ServeOptimized(w, r, []byte(html), ".html", false)
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
		emohttp.DisableContentEncodingForHTML = true
	}

	data.Init()

	err := goemo.InitSCSS(&godartsass.Options{
		LogEventHandler: func(e godartsass.LogEvent) {
			switch e.Type {
			case godartsass.LogEventTypeWarning:
				slog.Warn("sass: " + e.Message)
			case godartsass.LogEventTypeDeprecated:
				slog.Warn("sass deprecated: " + e.Message)
			case godartsass.LogEventTypeDebug:
				slog.Debug("sass debug: " + e.Message)
			}
		},
	})
	if err != nil {
		slog.Error("failed to start scss transpiler", "err", err.Error())
		os.Exit(1)
	}

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

	mux.HandleFunc("GET /{$}", handlePage(page.Index))
	mux.HandleFunc("GET /squirrels", handlePage(page.Squirrels))
	mux.HandleFunc("GET /webring", handlePage(page.Webring))
	mux.HandleFunc("GET /fav/anime", handlePage(page.FavAnime))
	mux.HandleFunc("GET /fav/games", handlePage(page.FavGames))

	// register assets

	mux.HandleFunc(
		"GET /cache/{file...}", emohttp.FileServerOptimized(
			os.DirFS("cache/public"),
		),
	)

	publicFS, err := fs.Sub(staticContent, "public")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	mux.HandleFunc("GET /{file...}", emohttp.FileServerOptimized(publicFS))

	// middleware

	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", util.GetGoVersion()) // hell yeah
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
