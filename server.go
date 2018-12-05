package main

import (
	"flag"
	"log"
	"net/http"
	//"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	fs := http.FileServer(http.Dir("frontend/"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(*addr, fs))
}
