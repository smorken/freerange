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

type Position struct {
	ID int64
	X  int32
	Y  int32
}

func (viewPort *LevelViewPort) GetVisibleSet(level *Level) []*Position {
	space := level.GetCollidingShapes(viewPort)
	positions := make([]*Position, space.Length())
	for i := 0; i < space.Length(); i++ {
		entity := space.Get(i).GetData().(Entity)
		positions[i].ID = entity.ID
		positions[i].X = entity.X
		positions[i].Y = entity.Y
	}
}

//GetDestroyList queries the level for the visible items in this view port.
//Any items that are not visible, either because they were destroyed server
//side or are just outside of the viewport's bounds, but exist in the
//viewPort's visibleEntities set are returned as the destroy list.
//This list of object are then destroyed client side.
//(set difference of level.visible - viewPort.visible)
func (viewPort *LevelViewPort) GetDestroyList([]*Position) []int64 {
	space := level.GetCollidingShapes(viewPort)
	for i := 0; i < space.Length(); i++ {
		entity := space.Get(i).GetData().(Entity)

	}
}

//GetCreateList queries the level for the visible items in this view port.
//Any items that are not currently in the viewPort's visibleEntities are added
//to the returned slice.(and stored in the viewPort's struct)
// These object are then created client side.
//(set difference of viewPort.visible - level.visible)
func (viewPort *LevelViewPort) GetCreateList(level *Level) []*Entity {

}

//GetMoveList returns a list of updated entity positions by comparing
//the x, y positions, and then updating the local viewport copy's positions
func (viewPort *LevelViewPort) GetMoveList(level *Level) []Position {

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
