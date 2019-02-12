package freerangeserver

import "testing"

func TestNewLevel(t *testing.T) {
	mockEntities := []Entity{CreateTestEntity(0), CreateTestEntity(0)}
	l := NewLevel(mockEntities)
	if l.nextID != BaseSharedEntityID+int64(len(mockEntities)) {
		t.Error("nextID not incremented")
	}
	if len(l.entities) != len(mockEntities) {
		t.Error("incorrect number of entities in collection")
	}
	if l.Space.Length() != len(mockEntities) {
		t.Error("incorrect number of entities in collision space")
	}
}

func TestSelect(t *testing.T) {
	mockEntities := []Entity{CreateTestEntity(0), CreateTestEntity(0), CreateTestEntity(0)}
	mockEntities[0].W = 10
	mockEntities[0].H = 10
	mockEntities[0].SetXY(10, 10)

	mockEntities[1].W = 10
	mockEntities[1].H = 10
	mockEntities[1].SetXY(20, 20)

	mockEntities[2].W = 10
	mockEntities[2].H = 10
	mockEntities[2].SetXY(30, 30)

	l := NewLevel(mockEntities)
	result := l.Select(0, 0, 10, 10)
	if len(result) != 0 {
		t.Error("expected 0 entities")
	}
	result = l.Select(10, 10, 10, 10)
	if len(result) != 1 {
		t.Error("expected 1 entities")
	}

	result = l.Select(15, 15, 10, 10)
	if len(result) != 2 {
		t.Error("expected 2 entities")
	}
	result = l.Select(20, 20, 20, 20)
	if len(result) != 3 {
		t.Error("expected 3 entities")
	}
}
