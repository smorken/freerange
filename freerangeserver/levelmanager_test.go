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
	mockLevelFactory := func(data []Entity) *Level {
		result := new(Level)
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

}
