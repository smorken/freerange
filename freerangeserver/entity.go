package freerangeserver

//Entity is a game object
type Entity struct {
	id             int64
	img            string
	tags           []string
	xposition      int
	yposition      int
	rotation       int
	xsize          int
	ysize          int
	static         bool
	clickable      bool
	parentEntityID int64
	cameraChild    bool
	cameraParent   bool
	zorder         int
}

type ClickAction struct {
	Action string //create, destroy
	id     int32  //for create, the id of an object def, for destroy the id of a live object

}

type create func(int) Entity

func makeCreateFunction(objectDefId int64, params map[string]string) func(int) Entity {

}

type ObjectDef struct {
	id      int64
	img     string
	onclick ClickAction
}

func getObjectDefs() string {
	objectdefs := `[
		{
			"id": 1
			"tags": ["actor"],
			"onclick": [
				{
					"action": "destroyAllByTag", 
					"params": {
						"parentEntityId": "$this",
						"xposition": "$ui0x"
					}
				},
				{
					"action": "create", 
					"params": {
						"objectDefId": 2
						"parentEntityId": "$this",
						"xposition": "$ui0x"
					}
				},
				{
					"action": "create", 
					"params": {
						"objectDefId": 3
						"parentEntityId": "$this",
						"xposition": "$ui1x"
					}
				},
			]	
		},
		{
			"id": 2
			"tags" ["ui", "left"]
			"onclick": [
				{
					"action": "update", 
					"params": {
						"entityid": "$parentEntityId"
						"xposition": "$ui0"
					}
				}
			]	
		},
		{
			"id": 3
			"tags" ["ui", "right"]
			"onclick": [
				{
					"action": "update", 
					"params": {
						"entityid": "$parentEntityId"
						"xposition": "$ui1"
					}
				}
			]	
		},
	]`
	return objectdefs
}

type CreateClickAction struct {
}
