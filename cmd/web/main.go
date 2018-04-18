package main

import (
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP listen address")
	htmlDir := flag.String("html-dir", "ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "ui/static", "Path to static assets")
	flag.Parse()

	app := &App{
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
