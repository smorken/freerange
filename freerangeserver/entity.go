package freerangeserver

//Entity is a game object
type Entity struct {
	ID             int64
	Img            string
	Tags           []string
	Xposition      int
	Yposition      int
	Rotation       int
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

func actorClick(client int64, clicked *Entity) func(level *Level) {
	return func(level *Level) {
		level.DestroyUIEntities(client)
		level.CreateUIEntity(client)
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

type CreateClickAction struct {
}
