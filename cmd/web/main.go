package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
    addr        string
    saticDir    string
}

type application struct {
    errorLog    *log.Logger
    infoLog     *log.Logger
}

func main() {
    var cfg config
    flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
    flag.StringVar(&cfg.saticDir, "static-dir", "./ui/static", "Path to static assets")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    srv := &http.Server{
        Addr: cfg.addr,
        Handler: app.routes(),
    }

    infoLog.Printf("Starting server on %s", cfg.addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
