package commands

import (
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func Initialize(arguments Arguments) error {
	pwd, err := wd.Get()
	if err != nil {
		return err
	}

	kd, makeErr := kaydir.Make(pwd)
	if makeErr != nil {
		return err
	}

	return index.Make(kd)
}
