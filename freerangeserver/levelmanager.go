package freerangeserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var levellock = sync.RWMutex{}

type ILevelManager interface {
	GetLevel(id int64, levelFactory LevelFactory, entityFactory EntityFactory) ILevel
	CloseLevel(level ILevel)
}
type LevelManager struct {
	directory     string
	levels        map[int64]ILevel
	levelRefCount map[int64]int32
}
type LevelFactory func(id int64, data []Entity) ILevel
type EntityFactory func(data map[string]interface{}) Entity

func NewLevelManager(directory string) *LevelManager {
	l := new(LevelManager)
	l.levels = make(map[int64]ILevel)
	l.levelRefCount = make(map[int64]int32)
	dir, err := filepath.Abs(directory)
	check(err)
	l.directory = dir
	return l
}

func (levelManager *LevelManager) getPath(id int64) string {
	return filepath.Join(levelManager.directory,
		fmt.Sprintf("%d.json", id))
}

func (levelManager *LevelManager) GetLevel(id int64, levelFactory LevelFactory,
	entityFactory EntityFactory) ILevel {
	levellock.Lock()
	defer levellock.Unlock()
	if lev, ok := levelManager.levels[id]; ok {
		return lev
	}
	dat, err := ioutil.ReadFile(levelManager.getPath(id))
	check(err)
	lev := levelFactory(id, deserializeEntities(dat, entityFactory))
	levelManager.levels[id] = lev
	levelManager.levelRefCount[id]++
	return lev
}
func (levelManager *LevelManager) CloseLevel(level *Level) {
	levellock.Lock()
	defer levellock.Unlock()
	levelManager.levelRefCount[level.ID]--
	if levelManager.levelRefCount[level.ID] == 0 {
		delete(levelManager.levelRefCount, level.ID)
		delete(levelManager.levels, level.ID)
	}
}
func deserializeEntities(data []byte, entityFactory EntityFactory) []Entity {

	result := []Entity{}
	deserialized := []interface{}{}
	err := json.Unmarshal(data, &deserialized)
	check(err)
	for _, item := range deserialized {
		values := item.(map[string]interface{})
		result = append(result, entityFactory(values))
	}
	return result

}
