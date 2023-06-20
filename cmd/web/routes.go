package main

import "net/http"

func (app *application) routes() *http.ServeMux {
    mux := http.NewServeMux()

    fileserver := http.FileServer(http.Dir("./ui/static"))
    mux.Handle("/static/", http.StripPrefix("/static", fileserver))

    return mux
}