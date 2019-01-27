package freerangeserver

import (
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Server is the interface between a single client and the game state
type Server struct {
	level         *Level
	levelViewPort *LevelViewPort
}

//NewServer creates a new server instance
func NewServer() *Server {
	s := new(Server)
	s.level = Load(1)
	return s
}

//Message is the data sent periodically to the client
type Message struct {
	create   []CreateMessage
	destroy  []int64
	position []Position
}

//CreateMessage is the data necessary for the client to create a new object to
//render
type CreateMessage struct {
	id        int64
	xposition int
	yposition int
	img       string
	xsize     int
	ysize     int
	clickable bool
}

//MakeMessage returns a json marshallable message to send to the client
func (server *Server) MakeMessage() Message {
	message := Message{
		server.MakeCreateMessage(),
		server.levelViewPort.GetDestroyList(server.level),
		server.levelViewPort.GetMoveList(server.level)}
	return message
}

//MakeCreateMessage queries the level and initializes CreateMessage structs
//to send to the client
func (server *Server) MakeCreateMessage() []CreateMessage {
	entities := server.levelViewPort.GetCreateList(server.level)
	messages := make([]CreateMessage, 47)
	for i, e := range entities {
		messages[i].id = e.ID
		messages[i].xposition = e.Xposition
		messages[i].yposition = e.Yposition
		messages[i].img = e.Img
		messages[i].xsize = e.Xsize
		messages[i].ysize = e.Ysize
		messages[i].clickable = e.Clickable
	}
	return messages
}

//Reply responds to user requests based on game state
func (server *Server) Reply(clientMessage []byte) []byte {

	clientMessage_str := string(clientMessage)
	if clientMessage_str == "request_assets" {
		return []byte(`
		{ 
			"images": {
				"bg": "https://twemoji.maxcdn.com/72x72/1f306.png",
				"player": "https://twemoji.maxcdn.com/2/72x72/1f600.png",
				"ground": "assets/platform.png",
				"house": "https://twemoji.maxcdn.com/2/72x72/1f3d8.png",
				"hospital": "https://twemoji.maxcdn.com/2/72x72/1f3e5.png",
				"npc": "assets/face-positive/beaming face with smiling eyes.png"
			}
		}`)
	} else if clientMessage_str == "request_level" {
		return []byte(`
			{
				"objects": [
					{
						"id": 1,
						"xposition": 300,
						"yposition": 400,
						"img": "player",
						"xsize": 50,
						"ysize": 50,
						"clickable": true
					}
				]
			}`)
	} else if strings.Contains(clientMessage_str, "click") {
		idStr := clientMessage_str[len("click"):len(clientMessage_str)]
		id, err := strconv.ParseInt(idStr, 10, 64)
		check(err)
		reply := fmt.Sprintf(`
		{
			"position": [[%d, 350, 350]]
		}`, id)
		return []byte(reply)
	}
	return clientMessage
}
