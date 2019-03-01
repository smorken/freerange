package freerangeserver

import "testing"

func TestNewGameContext(t *testing.T) {
	levelmanager := new(MockLevelManager)
	levelFactory := func(id int64, data []Entity) ILevel {
		return nil
	}
	entityFactory := func(data map[string]interface{}) Entity {
		return Entity{}
	}
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) ILevelViewPort {
		return nil
	}
	client := Client{1, "", 10, 10}
	g := NewGameContext(client, levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	if g.client.ID != 1 ||
		g.levelmanager != levelmanager ||
		g.levelFactory == nil ||
		g.entityFactory == nil ||
		g.levelViewPortFactory == nil {
		t.Error("values not assigned")
	}
}

func TestLoadLevel(t *testing.T) {
	levelmanager := new(MockLevelManager)
	getLevelCallCount := 0
	mockLevel := new(MockLevel)
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
	mockLevelViewPort := new(MockLevelViewPort)
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) ILevelViewPort {
		return mockLevelViewPort
	}
	client := Client{}
	g := NewGameContext(client, levelmanager, levelFactory, entityFactory, levelViewPortFactory)
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

func TestExit(t *testing.T) {
	levelmanager := new(MockLevelManager)
	mockLevel := new(MockLevel)
	levelmanager.mockGetLevel = func(int64, LevelFactory, EntityFactory) ILevel {
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
	mockLevelViewPort := new(MockLevelViewPort)
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) ILevelViewPort {
		return mockLevelViewPort
	}
	client := Client{}
	g := NewGameContext(client, levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	g.Exit()
	if closeLevelCallCount != 0 {
		t.Error("expected no call to close level")
	}
	g.LoadLevel(1)
	g.Exit()
	if closeLevelCallCount != 1 {
		t.Error("expected a single call to close level")
	}
}

func TestRefresh(t *testing.T) {
	levelmanager := new(MockLevelManager)
	mockLevel := new(MockLevel)

	levelmanager.mockGetLevel = func(int64, LevelFactory, EntityFactory) ILevel {
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
	mockRefreshCallCount := 0
	mockLevelViewPort := new(MockLevelViewPort)
	mockLevelViewPort.mockRefresh = func(ILevel) RefreshResult {
		mockRefreshCallCount++
		return RefreshResult{}
	}
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) ILevelViewPort {
		return mockLevelViewPort
	}
	client := Client{}
	g := NewGameContext(client, levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	g.LoadLevel(1)
	g.Refresh()
	if mockRefreshCallCount != 1 {
		t.Error("expected a call to level viewport instance")
	}
	g.Refresh()
	if mockRefreshCallCount != 2 {
		t.Error("expected a call to level viewport instance")
	}
}

func TestClickAction(t *testing.T) {
	levelmanager := new(MockLevelManager)
	mockLevel := new(MockLevel)
	entityClickCount := 0
	mockLevel.mockGetEntity = func(id int32) Entity {
		e := Entity{}
		e.clickAction = func(g *GameContext) {
			entityClickCount++
		}
		return e
	}
	levelmanager.mockGetLevel = func(int64, LevelFactory, EntityFactory) ILevel {
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
		e := Entity{}
		return e
	}
	mockRefreshCallCount := 0
	mockLevelViewPort := new(MockLevelViewPort)
	mockLevelViewPort.mockRefresh = func(ILevel) RefreshResult {
		mockRefreshCallCount++
		return RefreshResult{}
	}
	levelViewPortFactory := func(positionX int32, positionY int32, height int32, width int32) ILevelViewPort {
		return mockLevelViewPort
	}
	client := Client{}
	g := NewGameContext(client, levelmanager, levelFactory, entityFactory, levelViewPortFactory)
	g.LoadLevel(1)
	g.ClickAction(1)
	g.ClickAction(1)
	g.ClickAction(1)
	if entityClickCount != 3 {
		t.Error("expected 3 calls to clickaction")
	}

}
