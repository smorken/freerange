package freerangeserver

import "sync"

var lock = sync.RWMutex{}

//Level is a game state, at least 1 player is in the level
type Level struct {
	entities map[int64]*Entity
}

func Load(id int64) *Level {
	l := new(Level)
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
	delete(level.entities, id)
}

//AddEntity adds an entity to the level
func (level *Level) AddEntity(entity *Entity) {
	lock.Lock()
	defer lock.Unlock()
	id := int64(len(level.entities))
	entity.ID = int64(id)
	level.entities[id] = entity
}

func (level *Level) Move(entityId int64, direction string) {

}

func (level *Level) DestroyUIEntities(clientID int64) {

}

func (level *Level) AddUIEntity(clientID int64, entity *Entity) {

}

func (level *Level) SetCameraParent(clientID int64, entity *Entity) {

}
