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

type Server struct {
}

func NewServer() *Server {
	s := new(Server)
	return s
}

func (server *Server) InitializePhysics() {
	Box2dTutorial()
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
