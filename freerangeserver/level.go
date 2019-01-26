package freerangeserver

//Level is a game state, at least 1 player is in the level
type Level struct {
	entities map[int64]Entity
}

//Add entity adds an entity to the level
func (level *Level) AddEntity(entity *Entity) {

}

func (level *Level) DestroyUIEntities(clientID int64) {

}

func (level *Level) AddUIEntity(clientID int64, entity *Entity) {

}

func (level *Level) SetCameraParent(clientID int64, entity *Entity) {

}

func (level *Level) Move(entityId int64, direction string) {

}
