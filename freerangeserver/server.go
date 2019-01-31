package freerangeserver

import (
	"encoding/json"
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
	xposition int32
	yposition int32
	img       string
	xsize     int32
	ysize     int32
	clickable bool
}

//MakeCreateMessage queries the level and initializes CreateMessage structs
//to send to the client
func (server *Server) makeCreateMessage(entities []Entity) []createMessage {
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
	} else if clientMessage_str == "request_update" {
		RefreshResult := server.levelViewPort.Refresh(server.level)
		message := message{
			server.makeCreateMessage(RefreshResult.created),
			RefreshResult.destroyed,
			RefreshResult.moved}
		return serializeMessage(message)
	} else if strings.Contains(clientMessage_str, "click") {
		idStr := clientMessage_str[len("click"):len(clientMessage_str)]
		id, err := strconv.ParseInt(idStr, 10, 64)
		check(err)
		e := server.level.Read(id)
		e.clickAction(server.level, server.levelViewPort)
		entities := server.levelViewPort.GetUICreateList()
		message := message{
			server.makeCreateMessage(entities),
			server.levelViewPort.GetUIDestroyList(),
			server.levelViewPort.GetUIMoveList()}
		return serializeMessage(message)
	}
	return clientMessage
}
