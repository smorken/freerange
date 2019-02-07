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

type ILevel interface {
	Select(positionX int32, positionY int32, height int32, width int32) []Entity
}

//Level is a game state, at least 1 player is in the level
type Level struct {
	*resolv.Space
	nextID   int64
	entities map[int64]Entity
}

//NewLevel creates a new level instance, and the specified enties are added
func NewLevel(entities []Entity) *Level {
	l := new(Level)
	l.nextID = BaseSharedEntityID //this needs to be read in from any serialized level data
	l.Space = resolv.NewSpace()
	for _, e := range entities {
		l.AddEntity(e)
	}
	return l
}

//Select returns all level entities in the rectangle defined by the parameters
func (level *Level) Select(positionX int32, positionY int32, height int32, width int32) []Entity {
	rect := resolv.NewRectangle(positionX, positionY, width, height)
	selection := level.GetCollidingShapes(rect)
	result := []Entity{}
	for i := 0; i < selection.Length(); i++ {
		item := selection.Get(i).GetData().(Entity)
		result = append(result, item)
	}
	return result
}
func (level *Level) Read(id int64) Entity {
	lock.RLock()
	defer lock.RUnlock()
	return level.entities[id]
}

func (level *Level) DeleteEntity(id int64) {
	lock.Lock()
	defer lock.Unlock()
	level.Space.RemoveShape(level.entities[id])
	delete(level.entities, id)
}

//AddEntity adds an entity to the level
func (level *Level) AddEntity(entity Entity) {
	lock.Lock()
	defer lock.Unlock()
	entity.ID = level.nextID
	level.nextID++
	level.entities[entity.ID] = entity
	level.Space.AddShape(entity)
}

func (level *Level) Move(entityId int64, direction string) {

}
