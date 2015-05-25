package commands

import (
	"testing"

	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestInitializeMakesKayDir(t *testing.T) {
	tempdir.TempWd(func(working wd.WorkingDirectory) {
		Initialize(working)
		_, err := kaydir.Get(working)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}

func TestInitializeMakesIndexDir(t *testing.T) {
	tempdir.TempWd(func(working wd.WorkingDirectory) {
		Initialize(working)
		kd, _ := kaydir.Get(working)
		_, err := index.Get(kd)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
