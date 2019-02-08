package freerangeserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var levellock = sync.RWMutex{}

type LevelManager struct {
	directory string
	levels    map[int64]*Level
}

func NewLevelManager(directory string) *LevelManager {
	l := new(LevelManager)
	dir, err := filepath.Abs(directory)
	check(err)
	l.directory = dir
	return l
}

func (levelManager *LevelManager) GetPath(id int64) string {
	return filepath.Join(levelManager.directory,
		fmt.Sprintf("%d.json", id))
}
func (levelManager *LevelManager) GetLevel(id int64) *Level {
	levellock.Lock()
	defer levellock.Unlock()
	if lev, ok := levelManager.levels[id]; ok {
		return lev
	}
	dat, err := ioutil.ReadFile(levelManager.GetPath(id))
	check(err)
	lev := NewLevel(DeserializeLevel(dat))
	levelManager.levels[id] = lev
	return lev
}

func DeserializeLevel(data []byte) []Entity {

	result := []Entity{}
	deserialized := []interface{}{}
	json.Unmarshal(data, &deserialized)
	for i, item := range deserialized {
		values := item.(map[string]interface{})
		entity := NewEntity(
			values["img"].(string),
			values["tags"].([]string),
			values["xposition"].(int32),
			values["yposition"].(int32),
			values["rotation"].(float64),
			values["xsize"].(int32),
			values["ysize"].(int32))
		result = append(result, *entity)
	}
	return result

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
