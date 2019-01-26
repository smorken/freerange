package freerangeserver

//Level is a game state, at least 1 player is in the level
type Level struct {
	entities map[int64]Entity
}

func (level *Level) DestroyUIEntities(clientID int64) {

}

func (level *Level) CreateUIEntity(clientID int64, entity *Entity) {

}
