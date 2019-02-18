package freerangeserver

import (
	"math"

	"github.com/ByteArena/box2d"
)

const PixelsPerMetre float64 = 50.0

//PixelCoordToPhysicsCoord converts the specified pixel coordinate pair to physics coordinates
func PixelCoordToPhysicsCoord(x int32, y int32) (float64, float64) {
	return float64(x) / PixelsPerMetre, float64(y) / PixelsPerMetre
}

//PhysicsCoordToPixelCoord converts the specified physics coordinate pair to pixel coordinates
func PhysicsCoordToPixelCoord(x float64, y float64) (int32, int32) {
	return int32(math.Round(x * PixelsPerMetre)), int32(math.Round(y * PixelsPerMetre))
}

func AddEntityBody(world *box2d.B2World, entity *Entity) {
	if entity.Physics {
		bodyDef := box2d.NewB2BodyDef()
		px, py := PixelCoordToPhysicsCoord(entity.X, entity.Y)
		bodyDef.Position.Set(px, py)
		if entity.Static {
			bodyDef.Type = box2d.B2BodyType.B2_staticBody
		} else {
			bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
		}
		body := world.CreateBody(bodyDef)
		box := box2d.NewB2PolygonShape()
		box.SetAsBox(1.0, 1.0)
		fixture := body.CreateFixture(box, 1.0)
		fixture.SetFriction(0.3)
		entity.Body = body
		body.SetUserData(entity)
	}
}
