package freerangeserver

type Server struct {
}

func (server *Server) InitializePhysics() {
	Box2dTutorial()
}

//Reply responds to user requests based on game state
func (server *Server) Reply(clientMessage []byte) []byte {

	if string(clientMessage) == "request_assets" {
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
	} else if string(clientMessage) == "request_level" {
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
	}
	return clientMessage
}
