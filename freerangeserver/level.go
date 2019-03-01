package freerangeserver

import "sync"

import (
	"github.com/ByteArena/box2d"
	"github.com/SolarLune/resolv/resolv"
)

var lock = sync.RWMutex{}

//BaseSharedEntityID is the first value
//used in the shared (between clients) entity id space
//values smaller than this are reserved
const BaseSharedEntityID int32 = 10000

type ILevel interface {
	Select(positionX int32, positionY int32, height int32, width int32) []Entity
	GetEntity(id int32) Entity
	GetID() int64
}

//Level is a game state, at least 1 player is in the level
type Level struct {
	*resolv.Space
	id                 int64
	World              box2d.B2World
	contactListener    *ContactListener
	nextID             int32
	entities           map[int32]Entity
	intersectionMatrix map[int64]interface{}
}

//NewLevel creates a new level instance, and the specified enties are added
func NewLevel(id int64, entities []Entity) *Level {
	l := new(Level)
	l.id = id
	l.entities = make(map[int32]Entity)
	l.nextID = BaseSharedEntityID //this needs to be read in from any serialized level data
	l.Space = resolv.NewSpace()
	gravity := box2d.B2Vec2{X: 0.0, Y: -9.8}
	l.World = box2d.MakeB2World(gravity)
	l.contactListener = new(ContactListener)
	l.World.SetContactListener(l.contactListener)
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
		item := selection.Get(i).GetData().(int32)
		result = append(result, item)
	}
	return result
}

func (level *Level) GetEntity(id int32) Entity {
	lock.RLock()
	defer lock.RUnlock()
	return level.entities[id]
}

//DeleteEntity removes entitiy from level
//collection, and destroys collider and physics body
func (level *Level) DeleteEntity(id int32) {
	lock.Lock()
	defer lock.Unlock()
	entity := level.entities[id]
	level.Space.RemoveShape(entity.Rectangle)
	if entity.Body != nil {
		level.World.DestroyBody(entity.Body)
	}
	delete(level.entities, id)
}

//AddEntity adds an entity to the level
func (level *Level) AddEntity(entity Entity) {
	lock.Lock()
	defer lock.Unlock()
	entity.ID = level.nextID
	level.nextID++
	entity.Rectangle.SetData(entity.ID)
	level.Space.AddShape(entity.Rectangle)
	AddEntityBody(&level.World, &entity)
	level.entities[entity.ID] = entity
}

func (level *Level) UpdateIntersectionMatrix(entity Entity) {
	if entity.onIntersectEnter != nil {
		//entityIntersections, ok := level.intersectionMatrix[entity.ID]
		colliding := level.GetCollidingShapes(entity)
		for i := 0; i < colliding.Length(); i++ {
			//case1 the colliding shape is already present in the intersection map
			//case2 the colliding shape is not yet present add it and emit a on intersection enter event

		}
	}
}

func (level *Level) Step() {
	timeStep := 1.0 / 60.0
	velocityIterations := 6
	positionIterations := 2
	level.World.Step(timeStep, velocityIterations, positionIterations)
	for _, entity := range level.entities {
		level.UpdateIntersectionMatrix(entity)
	}
}
