package freerangeserver

import "testing"
import "os"

func TestNewLevelManager(t *testing.T) {
	d, e := os.Getwd()
	check(e)
	l := NewLevelManager(d)
	if l.directory != d {
		t.Error("directory not set")
	}

}

func TestGetLevel(t *testing.T) {
	d, _ := os.Getwd()
	l := NewLevelManager(d)
	path := "16.json"
	f, err := os.Create(path)
	check(err)
	defer f.Close()
	defer os.Remove(path)
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

	level := l.GetLevel(16, mockLevelFactory, mockEntityFactory)
	if level == nil {
		t.Error("level not created")
	}
	if mockLevelFactoryCallCount != 1 || mockEntityFactoryCallCount != 3 {
		t.Error("factory not called")
	}
	mockLevelFactoryCallCount = 0
	mockEntityFactoryCallCount = 0
	//on subsequent calls to GetLevel with the same level id, the already loaded level should be returned

	level = l.GetLevel(16, mockLevelFactory, mockEntityFactory)
	if mockLevelFactoryCallCount != 0 || mockEntityFactoryCallCount != 0 {
		t.Error("factory should not be called")
	}

}
