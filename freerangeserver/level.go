package freerangeserver

//Level is a game state, at least 1 player is in the level
type Level struct {
	entities map[int64]Entity
}

func (level *Level) ClickAction(id int64) {
	level.entities[id]
}
