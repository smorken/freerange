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
	d2 := []byte(`[
		{	
			"img": "bicycle",
			"tags": ["a"],
			"xposition": -5000,
			"yposition": 5000,
			"rotation": 0,
			"xsize": 10,
			"ysize": 10
		}
	]`)
	_, err = f.Write(d2)
	check(err)
	mockLevelFactory := func(data []Entity) *Level {

	}
	mockEntityFactory := func(data map[string]interface{}) *Entity {

	}
	l.GetLevel(16, mockLevelFactory, mockEntityFactory)
}
