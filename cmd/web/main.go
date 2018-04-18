package main

import (
	"log"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/carletes/lets-go/models"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP listen address")
	dbHost := flag.String("db-host", "localhost", "Database host name")
	dbName := flag.String("db-name", "snippetbox", "Database name")
	dbPassword := flag.String("db-password", "snippetbox", "Database password")
	dbPort := flag.Int("db-port", 5432, "Database port")
	dbTLS := flag.Bool("db-tls", false, "Use TLS in database connections")
	dbUser := flag.String("db-user", "snippetbox", "PostgreSQL uid")
	htmlDir := flag.String("html-dir", "ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "ui/static", "Path to static assets")

	flag.Parse()

	db, err := models.NewDatabase(*dbHost, *dbPort, *dbTLS, *dbUser, *dbPassword, *dbName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()

	app := &App{
		Database:  db,
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s", *addr)
	err = http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
