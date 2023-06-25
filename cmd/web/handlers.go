package main

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/BradPreston/go-portfolio/data"
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
    projects := []string{"one-with-the-riverbed","aarons-music-service","get-pokemon-data-from-api","kazoo48"}

    posts, err := internal.FetchPosts()
    if err != nil {
        app.serverError(w, err)
    }

    var templateData = struct{
        PreferredStack []string
        PreviousTech []string
        Posts []*internal.Post
        projects []string
    }{
        preferredStack,
        previousTech,
        posts,
        projects,
    }


    err = ts.ExecuteTemplate(w, "base", templateData)
    if err != nil {
        app.serverError(w, err)
        return
    }
}

func (app *application) work(w http.ResponseWriter, r *http.Request) {
    p := strings.Split(r.URL.Path, "/")[1]
    

    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/project.tmpl",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

    project, err := data.GetWorkData(p)
    if err != nil {
        app.notFound(w)
        app.errorLog.Println(err)
        return
    }
    
    err = ts.ExecuteTemplate(w, "base", project)
    if err != nil {
        app.serverError(w, err)
        return
    }
}
