package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
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

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
