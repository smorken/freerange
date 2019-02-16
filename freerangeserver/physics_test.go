package freerangeserver

import (
	"math"
	"testing"
)

func TestPixelCoordToPhysicsCoord(t *testing.T) {
	x, y := PixelCoordToPhysicsCoord(10, 20)
	if x != 10.0/PixelsPerMetre || y != 20.0/PixelsPerMetre {
		t.Error("unexpected output from PixelCoordToPhysicsCoord")
	}
}

func TestPhysicsCoordToPixelCoord(t *testing.T) {
	x, y := PhysicsCoordToPixelCoord(102.5, -230.2)
	if float64(x) != math.Round(102.5*PixelsPerMetre) ||
		float64(y) != math.Round(-230.2*PixelsPerMetre) {
		t.Error("unexpected output from PhysicsCoordToPixelCoord")
	}
}
