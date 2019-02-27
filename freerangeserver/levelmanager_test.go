package freerangeserver

import (
	"fmt"
	"os"
	"testing"
)

func TestNewLevelManager(t *testing.T) {
	d, e := os.Getwd()
	check(e)
	l := NewLevelManager(d)
	if l.directory != d {
		t.Error("directory not set")
	}

}

func createTestLevelData(id int64) string {
	path := fmt.Sprintf("%v.json", id)
	f, err := os.Create(path)
	defer f.Close()
	check(err)
	//epecting 1 call to entity factory per object
	d2 := []byte(`[
		{	
			"a": "1"
		},
		{	
			"b": "2"
		},
		{	
			"c": "3"
		}

	]`)
	_, err = f.Write(d2)
	check(err)
	return path
}

func TestGetLevel(t *testing.T) {
	levelID := int64(42)
	path := createTestLevelData(levelID)
	d, _ := os.Getwd()
	l := NewLevelManager(d)
	defer os.Remove(path)

	mockEntityFactoryCallCount := 0
	mockEntityFactory := func(data map[string]interface{}) Entity {
		entity := Entity{}
		mockEntityFactoryCallCount++
		return entity
	}
	mockLevelFactoryCallCount := 0
	mockLevelFactory := func(id int64, data []Entity) ILevel {
		result := new(MockLevel)
		mockLevelFactoryCallCount++
		return result
	}

	level := l.GetLevel(levelID, mockLevelFactory, mockEntityFactory)
	if level == nil {
		t.Error("level not created")
	}
	if mockLevelFactoryCallCount != 1 || mockEntityFactoryCallCount != 3 {
		t.Error("factory not called")
	}
	mockLevelFactoryCallCount = 0
	mockEntityFactoryCallCount = 0
	//on subsequent calls to GetLevel with the same level id, the already loaded level should be returned

	level = l.GetLevel(levelID, mockLevelFactory, mockEntityFactory)
	if mockLevelFactoryCallCount != 0 || mockEntityFactoryCallCount != 0 {
		t.Error("factory should not be called")
	}

}

func TestCloseLevel(t *testing.T) {
	path := createTestLevelData(1)
	d, _ := os.Getwd()
	l := NewLevelManager(d)
	defer os.Remove(path)
	mockLevelFactory := func(id int64, data []Entity) ILevel {
		result := new(MockLevel)
		result.mockGetID = func() int64 {
			return id
		}
		return result
	}
	mockEntityFactory := func(data map[string]interface{}) Entity {
		entity := Entity{}
		return entity
	}
	level := l.GetLevel(1, mockLevelFactory, mockEntityFactory)
	if l.levelRefCount[1] != 1 {
		t.Error("expected level reference increment")
	}
	if len(l.levels) != 1 {
		t.Error("expected single level")
	}

	l.CloseLevel(level)
	if len(l.levelRefCount) != 0 ||
		len(l.levels) != 0 {
		t.Error("expected empty level maps")
	}
}
