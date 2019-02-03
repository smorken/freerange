package freerangeserver

import "testing"

type MockLevel struct {
	mockselect func(int32, int32, int32, int32) []Entity
}

func (mock *MockLevel) Select(positionX int32, positionY int32, height int32, width int32) []Entity {
	return mock.mockselect(positionX, positionY, height, width)
}

func CreateTestEntity(id int64) Entity {
	e := *NewEntity("", []string{},
		-10, 20, 0, 0, 0, false, 50, 50, true, true, -1, true, false, 0)
	e.ID = id
	return e
}

//TestNewLevelViewPort tests assignment of the struct properties by the constructor
func TestNewLevelViewPort(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	if l.positionX != 5 || l.positionY != 10 || l.height != 100 || l.width != 101 {
		t.Error("NewLevelViewPort did not assign properties as expected")
	}
}

func TestRefresh(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{CreateTestEntity(1), CreateTestEntity(2), CreateTestEntity(3)}
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	result := l.Refresh(mockLevel)
	if len(result.created) != 3 ||
		result.created[0].ID != 1 ||
		result.created[1].ID != 2 ||
		result.created[2].ID != 3 {
		t.Error("expected 3 sequential values")
	}

	//add another entity
	mockEntities = append(mockEntities, CreateTestEntity(4))
	result = l.Refresh(mockLevel)
	if len(result.created) != 1 ||
		result.created[0].ID != 4 {
		t.Error("expected a single new value")
	}
	//simulate the level destroying an entity and or an entity leaving the
	//viewport rect and check the result in the viewport
	mockEntities = append(mockEntities[:1], mockEntities[2:]...)
	result = l.Refresh(mockLevel)
	if len(result.created) != 0 ||
		len(result.destroyed) != 1 ||
		result.destroyed[0] != 2 {
		t.Error("expected a single new value")
	}

	//simulate a moved entity
	mockEntities[0].X = 1
	mockEntities[0].Y = 2
	result = l.Refresh(mockLevel)
	if len(result.created) != 0 ||
		len(result.destroyed) != 0 ||
		len(result.moved) != 1 {
		t.Error("expected a move result")
	}
}
