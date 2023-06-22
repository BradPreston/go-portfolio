package main

import (
	"html/template"
	"net/http"

	"github.com/BradPreston/go-portfolio/internal"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/home.tmpl",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    preferredStack := []string{"NextJS","TypeScript","Node","Golang","SQL","NoSQL"}
    previousTech := []string{"Express","PHP","Laravel","Docker","MongoDB","PostgreSQL"}

    posts, err := internal.FetchPosts()
    app.serverError(w, err)

    var templateData = struct{
        PreferredStack []string
        PreviousTech []string
        Posts []*internal.Post
    }{
        preferredStack,
        previousTech,
        posts,
    }


    err = ts.ExecuteTemplate(w, "base", templateData)
    if err != nil {
        app.serverError(w, err)
        return
    }
}
