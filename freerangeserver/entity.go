package freerangeserver

//EntityType is the possible type of entities
type EntityType int

const(
	Player EntityType = 0
	Npc EntityType = 0
)
//Entity is a game object
type Entity struct {
	id int64
	img int
	xposition int
	yposition int
	rotation int
	static bool
	class EntityType
	clickable bool
	cameraChild bool
	cameraParent bool
	zorder int
}
