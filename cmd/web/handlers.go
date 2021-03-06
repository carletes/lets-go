package main

import (
	"net/http"
	"strconv"
)

func (app *App) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet ..."))
}

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.Database.LatestSnippets()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	app.RenderHTML(w, r, "home.page.html", &HTMLData{Snippets: snippets})
}

func (app *App) NewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the new snippet form"))
}

func (app *App) ShowSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Database.GetSnippet(id)
	if err != nil {
		app.ServerError(w, err)
		return
	}
	if snippet == nil {
		app.NotFound(w)
		return
	}

	app.RenderHTML(w, r, "show.page.html", &HTMLData{Snippet: snippet})
}
