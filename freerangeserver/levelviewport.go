package freerangeserver

//LevelViewPort is the subset of level data visible to a single client
type LevelViewPort struct {
	positionX           int32
	positionY           int32
	positionInvalidated bool
	height              int32
	width               int32
	//visible entities are the subset of level entities (shared between all clients) that are visible to the current client
	visibleEntities map[int64]Position
	//uiEntities are entities visible only to the current client
	uiIEntities          []*Entity
	addedUIEntities      []*Entity
	destroyedUIEntityIDs []int64
	nextUIEntityID       int64
	//cameraParent is the entity on which the view port is centered
	cameraParent *Entity
}

//NewLevelViewPort creates a new level view for a single client
func NewLevelViewPort(positionX int32, positionY int32, height int32, width int32) *LevelViewPort {
	l := new(LevelViewPort)
	l.positionX = positionX
	l.positionY = positionY
	l.height = height
	l.width = width
	l.nextUIEntityID = 1
	l.visibleEntities = make(map[int64]Position)
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
func (viewPort *LevelViewPort) Refresh(level ILevel) RefreshResult {
	visibleSet := viewPort.getVisibleSet(level)
	if viewPort.cameraParent != nil {
		viewPort.Move(viewPort.cameraParent.X, viewPort.cameraParent.Y)
	}
	return RefreshResult{
		append(viewPort.getCreateList(visibleSet), viewPort.getUICreateList()...),
		append(viewPort.getDestroyList(visibleSet), viewPort.getUIDestroyList()...),
		viewPort.getMoveList(visibleSet)}
}

func (viewPort *LevelViewPort) getVisibleSet(level ILevel) map[int64]Entity {

	selection := level.Select(viewPort.positionX, viewPort.positionY, viewPort.height, viewPort.width)
	result := map[int64]Entity{}
	for _, e := range selection {
		result[e.ID] = e
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
	if viewPort.positionInvalidated {
		//if the position of the viewport itself has changed, sent it as an update
		viewPort.positionInvalidated = false
		result = append(result, Position{0, viewPort.positionX, viewPort.positionY})
	}
	for id, currentPos := range visibleSet {
		newPosition := viewPort.visibleEntities[id]
		if currentPos.X != newPosition.X || currentPos.Y != newPosition.Y {
			result = append(result, newPosition)
			viewPort.visibleEntities[id] = newPosition
		}
	}
	return result
}

func (viewPort *LevelViewPort) getUIDestroyList() []int64 {
	result := []int64{}
	result = append(result, viewPort.destroyedUIEntityIDs...)
	viewPort.destroyedUIEntityIDs = nil
	return result
}

func (viewPort *LevelViewPort) getUICreateList() []Entity {
	result := []Entity{}
	for _, e := range viewPort.addedUIEntities {
		result = append(result, *e)
	}
	viewPort.addedUIEntities = nil
	return result
}

//DestroyUIEntities deletes all UI entities.
//The next call to Refresh will return the entity ids in the destroy list.
func (viewPort *LevelViewPort) DestroyUIEntities() {

	viewPort.nextUIEntityID = 1

	for _, e := range viewPort.uiIEntities {
		viewPort.destroyedUIEntityIDs = append(viewPort.destroyedUIEntityIDs, e.ID)
	}
	viewPort.uiIEntities = nil
}

//AddUIEntity adds the specified entity to the viewport's UI entity collection
//the ui entity will be emitted by the server to the client on the next update
func (viewPort *LevelViewPort) AddUIEntity(entity *Entity) {
	entity.ID = viewPort.nextUIEntityID
	viewPort.uiIEntities = append(viewPort.uiIEntities, entity)
	viewPort.addedUIEntities = append(viewPort.addedUIEntities, entity)
	viewPort.nextUIEntityID++

}

//SetCameraParent sets the specified entity as the camera parent, meaning the
//viewport's position will update according to the entities position
func (viewPort *LevelViewPort) SetCameraParent(entity *Entity) {
	viewPort.cameraParent = entity
}

//Move updates the position of the viewport
func (viewPort *LevelViewPort) Move(positionX int32, positionY int32) {
	if viewPort.positionX != positionX ||
		viewPort.positionY != positionY {
		viewPort.positionX = positionX
		viewPort.positionY = positionY
		viewPort.positionInvalidated = true
	}
}
