package freerangeserver

import (
	"github.com/SolarLune/resolv/resolv"
)

//Entity is a game object
type Entity struct {
	ID             int64
	Img            string
	Tags           []string
	Xposition      int
	Yposition      int
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
	collider       resolv.Rectangle
	clickAction    func(level *Level, levelviewport *LevelViewPort)
}

//NewEntity creates an entity with all fields specified by the function parameters
func NewEntity(Img string, Tags []string, Xposition int, Yposition int, Rotation float64,
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

func actorClick(clicked *Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		levelviewport.DestroyUIEntities()
		levelviewport.SetCameraParent(clicked)
		left := NewEntity("left arrow", []string{"ui", "left"},
			-10, 20, 0, 0, 0, false, 50, 50, true, true, clicked.ID, true, false, 0)
		left.clickAction = arrowClick(left)
		right := NewEntity("right arrow", []string{"ui", "right"},
			-10, 20, 0, 0, 0, false, 50, 50, true, true, clicked.ID, true, false, 0)
		right.clickAction = arrowClick(right)
		levelviewport.AddUIEntity(left)
		levelviewport.AddUIEntity(right)
	}
}

func arrowClick(entity *Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		level.Move(entity.ParentEntityID, entity.Tags[1])
	}
}
