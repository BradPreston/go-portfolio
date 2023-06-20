package main

import "flag"

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
}
