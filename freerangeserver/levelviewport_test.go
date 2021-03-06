package freerangeserver

import "testing"

func CreateTestEntity(id int32) Entity {
	e := MakeEntity("", []string{},
		-10, 20, 0, 10, 10)
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

func TestRefreshWithCameraParent(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 1000)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{CreateTestEntity(1), CreateTestEntity(2), CreateTestEntity(3)}
	mockEntities[1].X = 100
	mockEntities[1].Y = -100
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	result := l.Refresh(mockLevel)
	l.SetCameraParent(mockEntities[1])
	result = l.Refresh(mockLevel)
	if l.positionX != 100 && l.positionY != -100 {

		t.Error("expected the viewport position to change to the second mock entity")
	}
	if result.moved[0].ID != 0 {
		//the viewport itself is the 0th entity
		t.Error("expected a viewport move message")
	}
}
func TestRefreshWithMove(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{CreateTestEntity(1), CreateTestEntity(2), CreateTestEntity(3)}
	mockEntities[0].SetXY(0, 0)
	mockEntities[1].SetXY(0, 0)
	mockEntities[2].SetXY(0, 0)
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	//on the first call to refresh the entites are added to the viewport
	result := l.Refresh(mockLevel)

	//simulate a moved entity
	//verify an entity whose position did not change is not included in the update
	mockEntities[0].SetXY(1, 2)
	mockEntities[1].SetXY(0, 0)
	mockEntities[2].SetXY(-1, -2)

	result = l.Refresh(mockLevel)
	if len(result.moved) != 2 {
		t.Error("expected 2 move results")
	}
	resultMap := map[int32]Position{
		result.moved[0].ID: result.moved[0],
		result.moved[1].ID: result.moved[1]}

	if resultMap[1].X != 1 || resultMap[1].Y != 2 ||
		resultMap[3].X != -1 || resultMap[3].Y != -2 {
		t.Error("incorrect move results")
	}

}
func TestRefreshWithDestroy(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{CreateTestEntity(1), CreateTestEntity(2), CreateTestEntity(3)}
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	//on the first call to refresh the entites are added to the viewport
	result := l.Refresh(mockLevel)
	//simulate the level destroying an entity and or an entity leaving the
	//viewport rect and check the result in the viewport
	mockEntities = append(mockEntities[:1], mockEntities[2:]...)
	result = l.Refresh(mockLevel)
	if len(result.created) != 0 ||
		len(result.destroyed) != 1 ||
		result.destroyed[0] != 2 {
		t.Error("expected a single new value")
	}
}
func TestRefreshWithCreate(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{CreateTestEntity(1), CreateTestEntity(2), CreateTestEntity(3)}
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	result := l.Refresh(mockLevel)
	if len(result.created) != 3 {
		t.Error("expected 3 sequential values")
	}

	//add another entity
	mockEntities = append(mockEntities, CreateTestEntity(4))
	result = l.Refresh(mockLevel)
	if len(result.created) != 1 ||
		result.created[0].ID != 4 {
		t.Error("expected a single new value")
	}
}

func TestMove(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)

	mockLevel := new(MockLevel)
	mockEntities := []Entity{}
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}

	result := l.Refresh(mockLevel)
	//the very first refresh returns the viewport position
	if l.positionInvalidated != false ||
		l.positionX != 5 || l.positionY != 10 {
		t.Error("incorrect effect on levelviewport")
	}

	l.Move(5, 10)
	result = l.Refresh(mockLevel)
	//since the move was the same as the
	//initial position, the moved result should be empty
	if len(result.moved) != 0 ||
		l.positionInvalidated != false {
		t.Error("expected an empty move result")
	}

	l.Move(10, 5)
	if l.positionInvalidated != true ||
		l.positionX != 10 || l.positionY != 5 {
		t.Error("incorrect effect on levelviewport")
	}

	result = l.Refresh(mockLevel)
	if len(result.moved) != 1 ||
		result.moved[0].ID != 0 {
		t.Error("expected a move result")
	}

	//after the refresh, position invalidated should be reset
	if l.positionInvalidated != false {
		t.Error("positionInvalidated not reset")
	}
}

func TestAddDestroyUIEntities(t *testing.T) {
	l := NewLevelViewPort(5, 10, 100, 101)
	mockLevel := new(MockLevel)
	mockEntities := []Entity{}
	mockLevel.mockselect = func(int32, int32, int32, int32) []Entity {
		return mockEntities
	}
	l.Refresh(mockLevel)
	e1 := CreateTestEntity(1)
	e2 := CreateTestEntity(1)
	l.AddUIEntity(e1)
	l.AddUIEntity(e2)
	result := l.Refresh(mockLevel)
	if len(result.created) != 2 {
		t.Error("expected 2 ui entities")
	}
	result = l.Refresh(mockLevel)
	if len(result.created) != 0 {
		t.Error("expected 0 entities")
	}

	l.DestroyUIEntities()
	result = l.Refresh(mockLevel)
	if len(result.destroyed) != 2 {
		t.Error("expected 2 ui entities")
	}
}
