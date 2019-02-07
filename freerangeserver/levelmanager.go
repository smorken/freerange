package freerangeserver

import "sync"
import "encoding/json"

var levellock = sync.RWMutex{}

type LevelManager struct {
	levels map[int64]*Level
}

func (levelManager *LevelManager) GetLevel(id int64) *Level {
	levellock.Lock()
	defer levellock.Unlock()
	if lev, ok := levelManager.levels[id]; ok {
		return lev
	}
	lev := NewLevel(DeserializeLevel(""))
	levelManager.levels[id] = lev
	return lev
}

func DeserializeLevel(path string) []Entity {
	testJSON := []byte(`
	[{
		ID: 1
	}]`)
	deserialized := new(map[string]interface{})
	json.Unmarshal(testJSON, deserialized)

	ID := (*deserialized)["ID"].(int64)
		
}

//LoadAssets loads the assets needed to render the game state
func (levelManager *LevelManager) LoadAssets() []byte {
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
