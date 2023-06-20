package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
    addr        string
    saticDir    string
}

type application struct {}

func main() {
    var cfg config
    flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
    flag.StringVar(&cfg.saticDir, "static-dir", "./ui/static", "Path to static assets")
    flag.Parse()

    app := &application{}

    srv := &http.Server{
        Addr: cfg.addr,
        Handler: app.routes(),
    }

    err := srv.ListenAndServe()
    log.Fatal(err)
}
