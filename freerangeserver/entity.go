package freerangeserver

import (
	"github.com/SolarLune/resolv/resolv"
)

//Entity is a game object
type Entity struct {
	*resolv.Rectangle
	ID             int64
	Img            string
	Rotation       float64
	Speed          float64
	Jump           float64
	Fly            bool
	Xsize          int32
	Ysize          int32
	Static         bool
	Clickable      bool
	ParentEntityID int64
	CameraChild    bool
	CameraParent   bool
	Zorder         int32
	clickAction    func(level *Level, levelviewport *LevelViewPort)
}

//NewEntity creates an entity with all fields specified by the function parameters
func NewEntity(Img string, Tags []string, Xposition int32, Yposition int32, Rotation float64,
	Xsize int32, Ysize int32) *Entity {
	e := new(Entity)
	e.Rectangle = resolv.NewRectangle(Xposition, Yposition, Ysize, Xsize)
	e.Rectangle.SetData(e)
	e.Rectangle.SetTags(Tags...)
	e.Img = Img
	e.Speed = 0
	e.Jump = 0
	e.Fly = false
	e.Rotation = Rotation
	e.Static = true
	e.Clickable = false
	e.ParentEntityID = -1
	e.CameraChild = false
	e.CameraParent = false
	e.Zorder = 0
	e.ID = -1
	return e
}

func actorClick(clicked *Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		levelviewport.DestroyUIEntities()
		levelviewport.SetCameraParent(clicked)
		left := NewEntity("left arrow", []string{"ui", "left"},
			-10, 20, 0, 0, 0)
		left.clickAction = arrowClick(left)
		right := NewEntity("right arrow", []string{"ui", "right"},
			-10, 20, 0, 0, 0)
		right.clickAction = arrowClick(right)
		levelviewport.AddUIEntity(left)
		levelviewport.AddUIEntity(right)
	}
}

func arrowClick(entity *Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		level.Move(entity.ParentEntityID, entity.GetTags()[1])
	}
}
