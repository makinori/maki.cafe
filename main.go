package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed public
var staticContent embed.FS

func main() {
	_, inDev := os.LookupEnv("DEV")

	if inDev {
		http.Handle(
			"GET /", http.FileServer(http.Dir("public")),
		)
	} else {
		publicFs, err := fs.Sub(staticContent, "public")
		if err != nil {
			panic(err)
		}

		http.Handle(
			"GET /", http.FileServerFS(publicFs),
		)
	}

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	log.Println("listening at " + addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
