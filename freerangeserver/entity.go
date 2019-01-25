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

func getObjectDefs() string{
	objectdefs := `[
		{
			"id": 1
			"name": "actor",
			"onclick": [
				{
					"action": "create", 
					"params": {
						"id": 2,
						"parentEntityId"
					}
				}
			]	
		}
	]`
	return objectdefs
}

type CreateClickAction struct {
}
