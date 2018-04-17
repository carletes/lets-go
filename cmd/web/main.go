package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP listen address")
	htmlDir := flag.String("html-dir", "ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "ui/static", "Path to static assets")
	flag.Parse()

	app := &App{
		HTMLDir: *htmlDir,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/new", app.NewSnippet)

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
