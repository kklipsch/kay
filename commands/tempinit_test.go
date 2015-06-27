package commands

import (
	"testing"

	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func tempInit(action func(wd.WorkingDirectory, kaydir.KayDir)) error {
	var err error
	err = tempdir.TempWd(func(working wd.WorkingDirectory) {
		Initialize(working)
		kd, kerr := kaydir.Get(working)
		if kerr != nil {
			err = kerr
		} else {
			action(working, kd)
		}
	})
	return err
}

func TestCanGetIndexInInit(t *testing.T) {

	var idx index.Index
	tempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {
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
