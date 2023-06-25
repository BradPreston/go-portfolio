package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
    mux := chi.NewRouter()

    fileserver := http.FileServer(http.Dir("./ui/static"))
    mux.Handle("/static/*", http.StripPrefix("/static/", fileserver))

    mux.Get("/", app.home)
    mux.Get("/{project}", app.work)

    return mux
}
