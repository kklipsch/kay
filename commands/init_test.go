package commands

import (
	"testing"

	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestInitializeMakesKayDir(t *testing.T) {
	tempdir.In("ini-makes-kaydir", func(dir string) {
		working := wd.WorkingDirectory(dir)
		Initialize(nil, working)
		_, err := kaydir.Get(working)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestInitializeMakesIndexDir(t *testing.T) {
	tempdir.In("ini-makes-index", func(dir string) {
		working := wd.WorkingDirectory(dir)
		Initialize(nil, working)
		kd, _ := kaydir.Get(working)
		_, err := index.Get(kd)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
