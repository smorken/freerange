package freerangeserver

import "sync"

import (
	"github.com/SolarLune/resolv/resolv"
)

var lock = sync.RWMutex{}

//BaseSharedEntityID is the first value
//used in the shared (between clients) entity id space
//values smaller than this are reserved
const BaseSharedEntityID int64 = 10000

//Level is a game state, at least 1 player is in the level
type Level struct {
	nextID int64
	*resolv.Space
	entities map[int64]*Entity
}

func Load(id int64) *Level {
	l := new(Level)
	nextID = BaseSharedEntityID //this needs to be read in from any serialized level data
	l.Space = resolv.NewSpace()
	l.AddEntity(NewEntity("player", []string{"player"}, 200, 200, 0, 10, 10, false, 30, 30, false, true, -1, false, true, 0))
	return l
}

func (level *Level) Read(id int64) *Entity {
	lock.RLock()
	defer lock.RUnlock()
	return level.entities[id]
}

func (level *Level) Delete(id int64) {
	lock.Lock()
	defer lock.Unlock()
	level.Space.RemoveShape(level.entities[id])
	delete(level.entities, id)
}

//AddEntity adds an entity to the level
func (level *Level) AddEntity(entity *Entity) {
	lock.Lock()
	defer lock.Unlock()
	entity.ID = nextID
	nextID++
	level.entities[entity.ID] = entity
	level.Space.AddShape(entity)
}

func (level *Level) Move(entityId int64, direction string) {

}
