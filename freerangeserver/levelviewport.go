package freerangeserver

type LevelViewPort struct {
	positionX       int
	positionY       int
	viewPortHeight  int
	viewPortWidth   int
	visibleEntities map[int64]Position
}

type Position struct {
	ID int64
	X  int
	Y  int
}

//GetDestroyList queries the level for the visible items in this view port.
//Any items that are not visible, either because they were destroyed server
//side or are just outside of the viewport's bounds, but exist in the
//viewPort's visibleEntities set are returned as the destroy list.
//This list of object are then destroyed client side.
//(set difference of level.visible - viewPort.visible)
func (viewPort *LevelViewPort) GetDestroyList(level *Level) []int64 {

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
