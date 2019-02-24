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
	levelmanager.mockGetLevel = func(int64, LevelFactory, EntityFactory) ILevel {

	}
	levelmanager.mockCloseLevel = func(level ILevel) {

	}
	levelFactory := func(id int64, data []Entity) ILevel {

	}
	entityFactory := func(data map[string]interface{}) Entity {
		return Entity{}
	}
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) *LevelViewPort {
		return nil
	}
}
