package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
)

var (
	//go:embed public pages
	staticContent embed.FS

	//go:embed assets/template.html
	siteTemplateHTML string
	//go:embed assets/style.css
	siteStyleCSS string

	minifier *minify.M
)

func handleIndex(pageTemplate *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		siteTemplate := pageTemplate.Lookup("template.html")
		if siteTemplate == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server missing site template"))
			log.Println("missing site template")
			return
		}

		// execute page

		pageBuf := bytes.NewBuffer(nil)

		err := pageTemplate.Execute(pageBuf, pageData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to execute page template"))
			log.Println("failed to execute page template: " + err.Error())
			return
		}

		// execute site

		siteBuf := bytes.NewBuffer(nil)
		siteData := map[string]any{
			"Style": template.CSS(siteStyleCSS),
			"Body":  template.HTML(pageBuf.String()),
		}

		err = siteTemplate.Execute(siteBuf, siteData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to execute site template"))
			log.Println("failed to execute site template: " + err.Error())
			return
		}

		// minify and write

		err = minifier.Minify("text/html", w, siteBuf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to minify page"))
			log.Println("failed to minify page: " + err.Error())
			return
		}

		// w.Write(siteBuf.Bytes())
	}
}

func main() {
	_, inDev := os.LookupEnv("DEV")

	// register templates

	templates, err := template.ParseFS(staticContent, "pages/*.html")
	if err != nil {
		log.Fatalln("failed registering templates: " + err.Error())
	}

	_, err = templates.New("template.html").Parse(siteTemplateHTML)
	if err != nil {
		log.Fatalln("failed to register site template: " + err.Error())
	}

	var templateNames []string
	for _, t := range templates.Templates() {
		templateNames = append(templateNames, t.Name())
	}
	log.Println("parsed: " + strings.Join(templateNames, ", "))

	// register minifiers

	minifier = minify.New()
	minifier.Add("text/css", &css.Minifier{})
	minifier.Add("text/html", &html.Minifier{
		KeepDocumentTags:    true,
		KeepQuotes:          true,
		KeepDefaultAttrVals: true,
	})

	// register page handles

	http.HandleFunc("GET /{$}", handleIndex(templates.Lookup("index.html")))

	// register assets

	if inDev {
		http.Handle(
			"GET /", http.FileServer(http.Dir("public")),
		)
	} else {
		publicFs, err := fs.Sub(staticContent, "public")
		if err != nil {
			log.Fatalln(err)
		}

		http.Handle(
			"GET /", http.FileServerFS(publicFs),
		)
	}

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
		log.Fatalln(err)
	}
}
