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

    type link struct{
        Title   string
        Slug    string
    }

    preferredStack := []string{"NextJS","TypeScript","Node","Golang","SQL","NoSQL"}
    previousTech := []string{"Express","PHP","Laravel","Docker","MongoDB","PostgreSQL"}
    projects := []link{
        { "One With the Riverbed", "/one-with-the-riverbed" },
        { "Aaron's Music Service", "/aarons-music-service" },
        { "Get Pokemon Data From API", "/get-pokemon-data-from-api" },
        { "Kazoo48 Film Festival", "/kazoo48" },
    }
    hobbies := []string{"Reading on the couch","Playing video or board games","Playing guitar","Riding the stationary bike"}
    playlists := []link{
        {"Instrumental Madness","https://open.spotify.com/playlist/6yJYdOmWrsrV3F7rEEDxzT?si=2d738b5c33b24f16"},
        {"Future Funk","https://open.spotify.com/playlist/37i9dQZF1DX3LDIBRoaCDQ?si=c39b289e8d4a43e4"},
        {"Pokemon Lofi","https://open.spotify.com/playlist/59OrkYvGv0oM1KgPABU7nw?si=e09a56a6241e4cb1"},
    }
    contacts := []link{
        {"Email", "mailto:bap5393@gmail.com"},
        {"GitHub", "https://github.com/BradPreston"},
        {"LinkedIn", "https://linkedin.com/in/brad-preston"},
        {"Twitter", "https://twitter.com/braddoescoding"},
    }

    posts, err := internal.FetchPosts()
    if err != nil {
        app.serverError(w, err)
    }

    var templateData = struct{
        PreferredStack []string
        PreviousTech []string
        Posts []*internal.Post
        Projects []link
        Hobbies []string
        Playlists []link
        Contacts []link
    }{
        preferredStack,
        previousTech,
        posts,
        projects,
        hobbies,
        playlists,
        contacts,
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
