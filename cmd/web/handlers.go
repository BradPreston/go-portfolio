package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to my portfolio"))
}
