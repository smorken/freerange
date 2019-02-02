package freerangeserver

import "testing"

//TestNewLevelViewPort tests assignment of the struct properties by the constructor
func TestNewLevelViewPort(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	if l.Rectangle.X != 5 || l.Rectangle.Y != 10 || l.Rectangle.H != 100 || l.Rectangle.W != 101 {
		t.Error("NewLevelViewPort did not assign properties as expected")
	}
}
