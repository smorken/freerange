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
