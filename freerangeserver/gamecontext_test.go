package freerangeserver

import "testing"

type MockLevelManager struct {
	mockGetLevel   func(int64, LevelFactory, EntityFactory) ILevel
	mockCloseLevel func(level ILevel)
}

func (m *MockLevelManager) GetLevel(id int64, fac LevelFactory, entityFac EntityFactory) ILevel {
	return m.mockGetLevel(id, fac, entityFac)
}
func (m *MockLevelManager) CloseLevel(level ILevel) {
	m.mockCloseLevel(level)
}
func TestNewGameContext(t *testing.T) {
	levelmanager := new(MockLevelManager)
	levelFactory := func(id int64, data []Entity) ILevel {
		return nil
	}
	entityFactory := func(data map[string]interface{}) Entity {
		return Entity{}
	}
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) *LevelViewPort {
		return nil
	}
	g := NewGameContext(levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	if g.levelmanager != levelmanager ||
		g.levelFactory == nil ||
		g.entityFactory == nil ||
		g.levelViewPortFactory == nil {
		t.Error("entityFactory not assigned")
	}
}

func TestLoadLevel(t *testing.T) {
	levelmanager := new(MockLevelManager)
	getLevelCallCount := 0
	mockLevel := new(Level)
	levelmanager.mockGetLevel = func(int64, LevelFactory, EntityFactory) ILevel {
		getLevelCallCount++
		return mockLevel
	}
	closeLevelCallCount := 0
	levelmanager.mockCloseLevel = func(level ILevel) {
		closeLevelCallCount++

	}

	levelFactory := func(id int64, data []Entity) ILevel {
		return nil
	}
	entityFactory := func(data map[string]interface{}) Entity {
		return Entity{}
	}
	mockLevelViewPort := new(LevelViewPort)
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) *LevelViewPort {
		return mockLevelViewPort
	}
	g := NewGameContext(levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	g.LoadLevel(1)
	if g.level != mockLevel {
		t.Error("unexpected level")
	}
	if g.levelViewPort != mockLevelViewPort {
		t.Error("unexepected level view port")
	}
	if getLevelCallCount != 1 || closeLevelCallCount != 0 {
		t.Error("expected 1 get level call and no close level calls on first load")
	}
	g.LoadLevel(1)
	if getLevelCallCount != 2 || closeLevelCallCount != 1 {
		t.Error("expected 2 get level call and 1 close level calls after second load")
	}
}
