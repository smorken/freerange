package freerangeserver

//Entity is a game object
type Entity struct {
	ID             int64
	Img            string
	Tags           []string
	Xposition      float64
	Yposition      float64
	Rotation       float64
	Speed          float64
	Jump           float64
	Fly            bool
	Xsize          int
	Ysize          int
	Static         bool
	Clickable      bool
	ParentEntityID int64
	CameraChild    bool
	CameraParent   bool
	Zorder         int
	clickAction    func(level *Level)
}

//NewEntity creates an entity with all fields specified by the function parameters
func NewEntity(Img string, Tags []string, Xposition float64, Yposition float64, Rotation float64,
	Speed float64, Jump float64, Fly bool, Xsize int, Ysize int, Static bool, Clickable bool,
	ParentEntityID int64, CameraChild bool, CameraParent bool, Zorder int) *Entity {
	e := new(Entity)
	e.Img = Img
	e.Tags = Tags
	e.Xposition = Xposition
	e.Yposition = Yposition
	e.Speed = Speed
	e.Jump = Jump
	e.Fly = Fly
	e.Rotation = Rotation
	e.Xsize = Xsize
	e.Ysize = Ysize
	e.Static = Static
	e.Clickable = Clickable
	e.ParentEntityID = ParentEntityID
	e.CameraChild = CameraChild
	e.CameraParent = CameraParent
	e.Zorder = Zorder
	e.ID = -1
	return e
}

func actorClick(client int64, clicked *Entity) func(level *Level) {
	return func(level *Level) {
		level.DestroyUIEntities(client)
		level.SetCameraParent(client, clicked)
		left := NewEntity("left arrow", []string{"ui", "left"},
			-10, 20, 0, 0, 0, false, 50, 50, true, true, clicked.ID, true, false, 0)
		left.clickAction = arrowClick(left)
		right := NewEntity("right arrow", []string{"ui", "right"},
			-10, 20, 0, 0, 0, false, 50, 50, true, true, clicked.ID, true, false, 0)
		right.clickAction = arrowClick(right)
		level.AddUIEntity(client, left)
		level.AddUIEntity(client, right)
	}
}

func arrowClick(entity *Entity) func(level *Level) {
	return func(level *Level) {
		level.Move(entity.ParentEntityID, entity.Tags[1])
	}
}

func getObjectDefs() string {
	objectdefs := `[
        {
            "id": 1,
            "tags": ["actor"],
            "events": [
                {
					"trigger": "click",
                    "action": {
						"type": "destroyAllByTag", 
                    	"params": {
                        	"tag": "ui",
						}
					}
                },
                {
					"trigger": "click",
                    "action": {
						"type": "create",
						"params" {
							"objectDefId": 2,
							"parentEntityId": "$this",
							"xposition": "$ui0x"
						}
                    }
                },
                {
					"trigger": "click",
                    "action": {
						"type": "create",
						"params" {
							"objectDefId": 3,
							"parentEntityId": "$this",
							"xposition": "$ui1x"
						}
                    }
				},
				{
					"trigger": "proximityenter",
					"action": {
						"type": "create",
						"params" {
							"objectDefId": 4,
						}
					}
				}
			],
        },
        {
            "id": 2,
            "tags": ["ui", "left"],
            "onclick": [
                {
                    "action": "update", 
                    "params": {
                        "entityid": "$parentEntityId",
                        "xposition": "$ui0"
                    }
                }
            ]   
        },
        {
            "id": 3,
            "tags": ["ui", "right"],
            "onclick": [
                {
                    "action": "update", 
                    "params": {
                        "entityid": "$parentEntityId",
                        "xposition": "$ui1"
                    }
                }
            ]   
		},
		{
            "id": 4,
            "tags": ["ui", "enter"],
            "onclick": [
                {
                    "action": "update", 
                    "params": {
                        "entityid": "$parentEntityId",
                        "xposition": "$ui1"
                    }
                }
            ]   
        }
    ]`
	return objectdefs
}
