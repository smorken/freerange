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
	gamecontext *GameContext
}

//NewServer creates a new server instance
func NewServer(gamecontext *GameContext) *Server {
	s := new(Server)
	s.gamecontext = gamecontext
	return s
}

//CloseServer releases level handle
func (server *Server) CloseServer() {

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
		messages[i].xposition = e.X
		messages[i].yposition = e.Y
		messages[i].img = e.Img
		messages[i].xsize = e.W
		messages[i].ysize = e.H
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

	clientMessageStr := string(clientMessage)
	if clientMessageStr == "request_assets" {
		return server.gamecontext.LoadAssets()
	} else if clientMessageStr == "request_update" {
		RefreshResult := server.levelViewPort.Refresh(server.level)
		message := message{
			server.makeCreateMessage(RefreshResult.created),
			RefreshResult.destroyed,
			RefreshResult.moved}
		return serializeMessage(message)
	} else if strings.Contains(clientMessageStr, "click") {
		idStr := clientMessageStr[len("click"):len(clientMessageStr)]
		id, err := strconv.ParseInt(idStr, 10, 64)
		check(err)
		e := server.level.GetEntity(id)
		e.clickAction(server.level, server.levelViewPort)
		return []byte("click")
	} else {
		return []byte("error")
	}
}
