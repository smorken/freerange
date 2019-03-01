package freerangeserver

import "testing"

func TestNewLevel(t *testing.T) {

	physicsBodyEntity1 := CreateTestEntity(0)
	physicsBodyEntity1.Physics = true
	physicsBodyEntity1.Static = false

	physicsBodyEntity2 := CreateTestEntity(0)
	physicsBodyEntity2.Physics = true
	physicsBodyEntity2.Static = true

	mockEntities := []Entity{physicsBodyEntity1, CreateTestEntity(0), physicsBodyEntity2}
	l := NewLevel(1, mockEntities)
	if l.nextID != BaseSharedEntityID+int32(len(mockEntities)) {
		t.Error("nextID not incremented")
	}
	if len(l.entities) != len(mockEntities) {
		t.Error("incorrect number of entities in collection")
	}
	if l.Space.Length() != len(mockEntities) {
		t.Error("incorrect number of entities in collision space")
	}
	if l.World.GetBodyCount() != 2 {
		t.Error("incorrect number of physics bodies")
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

	l := NewLevel(1, mockEntities)
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
	result = l.Select(0, 0, 50, 50)
	if len(result) != 3 {
		t.Error("expected 3 entities")
	}
}

func TestDeleteEntity(t *testing.T) {
	physicsBodyEntity := CreateTestEntity(0)
	physicsBodyEntity.Physics = true
	mockEntities := []Entity{physicsBodyEntity, CreateTestEntity(0), physicsBodyEntity, CreateTestEntity(0)}
	l := NewLevel(1, mockEntities)
	deleteList := []int32{}
	for k := range l.entities {
		deleteList = append(deleteList, k)
	}

	for _, d := range deleteList {
		l.DeleteEntity(d)
	}
	if len(l.entities) != 0 {
		t.Error("unexpected number of entities")
	}
	if l.Space.Length() != 0 {
		t.Error("incorrect number of entities in collision space")
	}
	if l.World.GetBodyCount() != 0 {
		t.Error("incorrect number of physics bodies")
	}
}

func TestGetEntity(t *testing.T) {
	mockEntities := []Entity{CreateTestEntity(0), CreateTestEntity(0), CreateTestEntity(0)}
	l := NewLevel(1, mockEntities)
	selectList := []int32{}
	for k := range l.entities {
		selectList = append(selectList, k)
	}
	for _, d := range selectList {
		e := l.GetEntity(d)
		if e.ID != d {
			t.Error("expected entity not returned")
		}
	}

}
