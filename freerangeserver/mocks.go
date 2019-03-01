package freerangeserver

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

type MockLevelViewPort struct {
	mockRefresh func(level ILevel) RefreshResult
}

func (m *MockLevelViewPort) Refresh(level ILevel) RefreshResult {
	return m.mockRefresh(level)
}

type MockLevel struct {
	mockselect    func(int32, int32, int32, int32) []Entity
	mockGetEntity func(id int32) Entity
	mockGetID     func() int64
}

func (mock *MockLevel) Select(positionX int32, positionY int32, height int32, width int32) []Entity {
	return mock.mockselect(positionX, positionY, height, width)
}

func (mock *MockLevel) GetEntity(id int32) Entity {
	return mock.mockGetEntity(id)
}

func (mock *MockLevel) GetID() int64 {
	return mock.mockGetID()
}
