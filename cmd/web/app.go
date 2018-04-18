package main

import "github.com/carletes/lets-go/models"

// App is the application state.
type App struct {
	HTMLDir   string
	StaticDir string
	Database  *models.Database
}
