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
		{	}
	]`)
	_, err = f.Write(d2)
	check(err)

	l.GetLevel(16)
}
