package main

/*
import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
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
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func loadTemplate() *template.Template {
	//var homeTemplate = template.ParseFiles.Must(template.New("").Parse()
	websockets, err := template.ParseFiles(
		"frontend/websockets.html")
	if err != nil {
		log.Println("template:", err)
		return nil
	}
	return websockets
}

var homeTemplate = loadTemplate()

//PageData is a structure to store template data
type PageData struct {
	SocketAddress string
}

func home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		SocketAddress: "ws://" + r.Host + "/echo",
	}
	err := homeTemplate.Execute(w, data)
	if err != nil {
		log.Println("template execute:", err)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	fs := http.FileServer(http.Dir("frontend/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fs))
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
*/
