package freerangeserver

type GameContext struct {
	levelmanager  *LevelManager
	level         *Level
	levelViewPort *LevelViewPort
}

//LoadAssets loads the assets needed to render the game state
func (gamecontext *GameContext) LoadAssets() []byte {
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
}
