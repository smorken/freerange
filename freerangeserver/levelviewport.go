package freerangeserver

import (
	"github.com/SolarLune/resolv/resolv"
)

type LevelViewPort struct {
	*resolv.Rectangle
	visibleEntities map[int64]Position
}

func NewLevelViewPort(positionX int32, positionY int32, height int32, width int32) *LevelViewPort {
	l := new(LevelViewPort)
	l.Rectangle = resolv.NewRectangle(positionX, positionY, width, height)
	l.Rectangle.SetData(l)
	return l
}

//Position is the x, y coordinate for the specified entity id
type Position struct {
	ID int64
	X  int32
	Y  int32
}

//RefreshResult is data passed to client on viewport syncronization with level
type RefreshResult struct {
	created   []Entity
	destroyed []int64
	moved     []Position
}

//Refresh updates this viewport instance according to the level state,
// and returns a result used to syncronize the client
func (viewPort *LevelViewPort) Refresh(level *Level) RefreshResult {
	visibleSet := viewPort.getVisibleSet(level)
	return RefreshResult{
		viewPort.getCreateList(visibleSet),
		viewPort.getDestroyList(visibleSet),
		viewPort.getMoveList(visibleSet)}
}

func (viewPort *LevelViewPort) getVisibleSet(level *Level) map[int64]Entity {
	space := level.GetCollidingShapes(viewPort)
	result := map[int64]Entity{}
	for i := 0; i < space.Length(); i++ {
		entity := space.Get(i).GetData().(Entity)
		result[entity.ID] = entity
	}
	return result
}

//getDestroyList queries the level for the visible items in this view port.
//Any items that are not visible, either because they were destroyed server
//side or are just outside of the viewport's bounds, but exist in the
//viewPort's visibleEntities set are returned as the destroy list.
//This list of object are then destroyed client side.
//(set difference of viewPort.visible - level.visible )
func (viewPort *LevelViewPort) getDestroyList(visibleSet map[int64]Entity) []int64 {

	result := []int64{}

	for id := range viewPort.visibleEntities {
		if _, ok := visibleSet[id]; !ok {
			result = append(result, id)
		}
	}
	for _, id := range result {
		//if the viewport's entity id is not found in the level's visible set, delete it
		delete(viewPort.visibleEntities, id)
	}
	return result
}

//getCreateList queries the level for the visible items in this view port.
//Any items that are not currently in the viewPort's visibleEntities are added
//to the returned slice.(and stored in the viewPort's struct)
// These object are then created client side.
//(set difference of level.visible - viewPort.visible )
func (viewPort *LevelViewPort) getCreateList(visibleSet map[int64]Entity) []Entity {

	result := []Entity{}

	for id, entity := range visibleSet {
		if _, ok := viewPort.visibleEntities[id]; !ok {
			result = append(result, entity)
			//if the level's entity id is not found in the viewport's set, add it to the viewport
			viewPort.visibleEntities[id] = Position{entity.ID, entity.X, entity.Y}
		}
	}
	return result
}

//getMoveList returns a list of updated entity positions by comparing
//the x, y positions, and then updating the local viewport copy's positions
//should be run after delete and create routines
func (viewPort *LevelViewPort) getMoveList(visibleSet map[int64]Entity) []Position {

	result := []Position{}
	for id, currentPos := range visibleSet {
		newPosition := viewPort.visibleEntities[id]
		if currentPos.X != newPosition.X || currentPos.Y != newPosition.Y {
			result = append(result, newPosition)
			viewPort.visibleEntities[id] = newPosition
		}
	}
	return result
}

func (viewPort *LevelViewPort) GetUIDestroyList() []int64 {

}

func (viewPort *LevelViewPort) GetUICreateList() []*Entity {

}

func (viewPort *LevelViewPort) GetUIMoveList() []Position {

}

func (viewPort *LevelViewPort) DestroyUIEntities() {

}

func (viewPort *LevelViewPort) AddUIEntity(entity *Entity) {

}

func (viewPort *LevelViewPort) SetCameraParent(entity *Entity) {

}
