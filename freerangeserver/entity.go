package freerangeserver

//Entity is a game object
type Entity struct {
	id             int64
	img            int
	xposition      int
	yposition      int
	rotation       int
	xsize          int
	ysize          int
	static         bool
	clickable      bool
	parentEntityId int64
	cameraChild    bool
	cameraParent   bool
	zorder         int
}

type ClickAction struct {
	entityDelta Entity
	destroy     bool
	create      bool
}

type CreateClickAction struct {
}
