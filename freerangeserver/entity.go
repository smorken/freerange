package freerangeserver

import (
	"github.com/ByteArena/box2d"
	"github.com/SolarLune/resolv/resolv"
)

//Entity is a game object
type Entity struct {
	*resolv.Rectangle
	Body             *box2d.B2Body
	ID               int64
	Img              string
	Rotation         float64
	Speed            float64
	Jump             float64
	Fly              bool
	Static           bool
	Physics          bool
	Clickable        bool
	ParentEntityID   int64
	CameraChild      bool
	CameraParent     bool
	Zorder           int32
	intersecting     []int64
	clickAction      func(gameContext *GameContext)
	onIntersectEnter func(gameContext *GameContext, otherEntity Entity)
	onIntersectLeave func(gameContext *GameContext, otherEntity Entity)
	onCollision      func(gameContext *GameContext, otherEntity Entity)
}

//MakeEntity creates an entity with all fields specified by the function parameters
func MakeEntity(Img string, Tags []string, Xposition int32, Yposition int32, Rotation float64,
	Xsize int32, Ysize int32) Entity {
	e := Entity{}
	e.Rectangle = resolv.NewRectangle(Xposition, Yposition, Ysize, Xsize)
	//e.Rectangle.SetData(e)
	//e.Rectangle.SetTags(Tags...)
	e.Img = Img
	e.Speed = 0
	e.Jump = 0
	e.Fly = false
	e.Rotation = Rotation
	e.Static = true
	e.Physics = false
	e.Clickable = false
	e.ParentEntityID = -1
	e.CameraChild = false
	e.CameraParent = false
	e.Zorder = 0
	e.ID = -1
	return e
}

func DeserializeEntity(values map[string]interface{}) Entity {
	tagI := values["tags"].([]interface{})
	tagStr := []string{}
	for _, t := range tagI {
		tagStr = append(tagStr, t.(string))
	}

	entity := MakeEntity(
		values["img"].(string),
		tagStr,
		int32(values["xposition"].(float64)),
		int32(values["yposition"].(float64)),
		values["rotation"].(float64),
		int32(values["xsize"].(float64)),
		int32(values["ysize"].(float64)))
	return entity
}

/* func actorClick(clicked Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		levelviewport.DestroyUIEntities()
		levelviewport.SetCameraParent(clicked)
		left := MakeEntity("left arrow", []string{"ui", "left"},
			-10, 20, 0, 0, 0)
		left.clickAction = arrowClick(left)
		right := MakeEntity("right arrow", []string{"ui", "right"},
			-10, 20, 0, 0, 0)
		right.clickAction = arrowClick(right)
		levelviewport.AddUIEntity(left)
		levelviewport.AddUIEntity(right)
	}
}

func arrowClick(entity Entity) func(level *Level, levelviewport *LevelViewPort) {
	return func(level *Level, levelviewport *LevelViewPort) {
		parent := level.GetEntity(entity.ParentEntityID)
		direction := entity.GetTags()[1]
		impulse := box2d.B2Vec2{0.0, 0.0}
		if direction == "left" {
			impulse.X = -parent.Speed
		} else if direction == "right" {
			impulse.X = parent.Speed
		} else if direction == "up" {
			impulse.Y = parent.Speed
		} else if direction == "down" {
			impulse.Y = -parent.Speed
		}

		parent.Body.ApplyLinearImpulseToCenter(impulse, true)

	}
} */
