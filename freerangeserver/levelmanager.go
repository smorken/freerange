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
type LevelFactory func(data []Entity) *Level
type EntityFactory func(data map[string]interface{}) *Entity

func NewLevelManager(directory string) *LevelManager {
	l := new(LevelManager)
	l.levels = make(map[int64]*Level)
	dir, err := filepath.Abs(directory)
	check(err)
	l.directory = dir
	return l
}

func (levelManager *LevelManager) getPath(id int64) string {
	return filepath.Join(levelManager.directory,
		fmt.Sprintf("%d.json", id))
}

func (levelManager *LevelManager) GetLevel(id int64, factory LevelFactory,
	entityFactory EntityFactory) *Level {
	levellock.Lock()
	defer levellock.Unlock()
	if lev, ok := levelManager.levels[id]; ok {
		return lev
	}
	dat, err := ioutil.ReadFile(levelManager.getPath(id))
	check(err)
	lev := factory(deserializeLevel(dat, entityFactory))
	levelManager.levels[id] = lev
	return lev
}

func deserializeLevel(data []byte, entityFactory EntityFactory) []Entity {

	result := []Entity{}
	deserialized := []interface{}{}
	err := json.Unmarshal(data, &deserialized)
	check(err)
	for _, item := range deserialized {
		values := item.(map[string]interface{})
		tagI := values["tags"].([]interface{})
		tagStr := []string{}
		for _, t := range tagI {
			tagStr = append(tagStr, t.(string))
		}

		entity := NewEntity(
			values["img"].(string),
			tagStr,
			int32(values["xposition"].(float64)),
			int32(values["yposition"].(float64)),
			values["rotation"].(float64),
			int32(values["xsize"].(float64)),
			int32(values["ysize"].(float64)))
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
