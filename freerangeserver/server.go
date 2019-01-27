package freerangeserver

import (
	"encoding/json"
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
	levelmanager  *LevelManager
	level         *Level
	levelViewPort *LevelViewPort
}

//NewServer creates a new server instance
func NewServer(levelmanager *LevelManager) *Server {
	s := new(Server)
	s.level = levelmanager.GetLevel(1)
	return s
}

//message is the data sent periodically to the client
type message struct {
	create   []createMessage
	destroy  []int64
	position []Position
}

//createMessage is the data necessary for the client to create a new object to
//render
type createMessage struct {
	id        int64
	xposition int
	yposition int
	img       string
	xsize     int
	ysize     int
	clickable bool
}

//MakeMessage returns a json marshallable message to send to the client
func (server *Server) makeMessage() message {
	message := message{
		server.makeCreateMessage(),
		server.levelViewPort.GetDestroyList(server.level),
		server.levelViewPort.GetMoveList(server.level)}
	return message
}

//MakeCreateMessage queries the level and initializes CreateMessage structs
//to send to the client
func (server *Server) makeCreateMessage() []createMessage {
	entities := server.levelViewPort.GetCreateList(server.level)
	messages := make([]createMessage, 47)
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

func serializeMessage(message message) []byte {
	j, e := json.Marshal(message)
	check(e)
	return j
}

//Reply responds to user requests based on game state
func (server *Server) Reply(clientMessage []byte) []byte {

	clientMessage_str := string(clientMessage)
	if clientMessage_str == "request_assets" {
		return server.levelmanager.LoadAssets()
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
