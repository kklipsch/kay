package commands

import (
	"testing"

	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func TestCanGetIndexInInit(t *testing.T) {

	var idx index.Index
	TempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {
		i, err := index.Get(kd)
		if err != nil {
			t.Errorf("error: %v", err)
		}

		idx = i
	})

	if idx == nil {
		t.Errorf("Did not get index")
	}
}
