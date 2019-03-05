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
	lock.RLock()
	defer lock.RUnlock()
	rect := resolv.NewRectangle(positionX, positionY, width, height)
	selection := level.GetCollidingShapes(rect)
	result := []Entity{}
	for i := 0; i < selection.Length(); i++ {
		id := selection.Get(i).GetData().(int32)
		result = append(result, level.entities[id])
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

//BuildIntersectionMatrix creates a matrix of entity pairs that are
//intersecting using a composite key based on the entity ids
func (level *Level) BuildIntersectionMatrix() map[int64]interface{} {
	lock.RLock()
	defer lock.RUnlock()
	intersections := map[int64]interface{}{}
	for _, entity := range level.entities {

		if entity.onIntersectEnter != nil {
			//entityIntersections, ok := level.intersectionMatrix[entity.ID]
			colliding := level.GetCollidingShapes(entity)
			for i := 0; i < colliding.Length(); i++ {
				otherEntityID := colliding.Get(i).GetData().(int32)
				if otherEntityID == entity.ID { //get colliding shapes returns the calling shape
					continue
				}
				key := int64(0)
				//always set the entity with the higher id value as the upper 32 bits
				//of the key, so that each entity pair cannot have more than 1 key
				if entity.ID >= otherEntityID {
					key = (int64(entity.ID) << 32) + int64(otherEntityID)
				} else {
					key = (int64(otherEntityID) << 32) + int64(entity.ID)
				}
				if _, ok := intersections[key]; !ok {
					intersections[key] = nil
				}
			}
		}
	}
	return intersections
}

//EmitIntersectionEvents compairs the entity pair intersection matrix for this
//step and the previous step, and emits intersect enter events for
//intersections that did not appear in the previous timestep, and intersect
// exit events for intersections that appeared in the previous timestep but not in the current timestep
func (level *Level) EmitIntersectionEvents(gameContext *GameContext, intersections map[int64]interface{}) {
	//for each value that occurs in the current intersection matrix
	//if the value does not yet occur in the level's matrix emit a onIntersectionEnter event
	for k := range intersections {
		if _, ok := level.intersectionMatrix[k]; !ok {
			e1 := level.entities[int32(k>>32)]
			e2 := level.entities[int32(k)]
			if e1.onIntersectEnter != nil {
				e1.onIntersectEnter(gameContext, e2)
			}
			if e2.onIntersectEnter != nil {
				e2.onIntersectEnter(gameContext, e1)
			}
		}
	}
	//for each value that occurs in the level's matrix, but not in the updated matrix
	//emit an onintersection leave event
	for k := range level.intersectionMatrix {
		if _, ok := intersections[k]; !ok {
			e1 := level.entities[int32(k>>32)]
			e2 := level.entities[int32(k)]
			if e1.onIntersectExit != nil {
				e1.onIntersectExit(gameContext, e2)
			}
			if e2.onIntersectExit != nil {
				e2.onIntersectExit(gameContext, e1)
			}
		}
	}
}

//Step advances the level physics state by one step, and emits time based events
func (level *Level) Step(gameContext *GameContext) {
	timeStep := 1.0 / 60.0
	velocityIterations := 6
	positionIterations := 2
	level.World.Step(timeStep, velocityIterations, positionIterations)

	mat := level.BuildIntersectionMatrix()
	level.EmitIntersectionEvents(gameContext, mat)
	//overwrite the intersection matrix
	level.intersectionMatrix = mat
}
