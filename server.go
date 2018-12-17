package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func clientLoop(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, []byte{'a'})
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	fs := http.FileServer(http.Dir("frontend/"))
	http.Handle("/", fs)
	http.HandleFunc("/client", clientLoop)
	log.Fatal(http.ListenAndServe(*addr, fs))
}
