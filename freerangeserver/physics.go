package freerangeserver

import (
	"fmt"

	"github.com/ByteArena/box2d"
)

//Box2dTutorial is from http://box2d.org/manual.pdf
func Box2dTutorial() {
	gravity := box2d.B2Vec2{X: 0.0, Y: -9.8}
	world := box2d.MakeB2World(gravity)

	ground := box2d.NewB2BodyDef()
	ground.Position.Set(0, -10)
	groundBody := world.CreateBody(ground)
	groundBox := box2d.NewB2PolygonShape()
	groundBox.SetAsBox(50.0, 10.0)
	groundBody.CreateFixture(groundBox, 0.0)

	bodyDef := box2d.NewB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(0.0, 4.0)
	body := world.CreateBody(bodyDef)
	dynamicBox := box2d.NewB2PolygonShape()
	dynamicBox.SetAsBox(1.0, 1.0)
	dynamicFixture := body.CreateFixture(dynamicBox, 1.0)
	dynamicFixture.SetFriction(0.3)

	timeStep := 1.0 / 60.0
	velocityIterations := 6
	positionIterations := 2
	for i := 0; i < 60; i++ {
		world.Step(timeStep, velocityIterations, positionIterations)
		position := body.GetPosition()
		angle := body.GetAngle()
		fmt.Printf("%4.2f %4.2f %4.2f\n", position.X, position.Y, angle)
	}

}
