package commands

import (
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TempInit(action func(wd.WorkingDirectory, kaydir.KayDir)) error {
	var err error
	err = tempdir.TempWd(func(working wd.WorkingDirectory) {
		Initialize(Arguments{}, working)
		kd, kerr := kaydir.Get(working)
		if kerr != nil {
			err = kerr
		} else {
			action(working, kd)
		}
	})
	return err
}
