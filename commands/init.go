package commands

import (
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func Initialize(arguments Arguments, working wd.WorkingDirectory) error {
	kd, makeErr := kaydir.Make(working)
	if makeErr != nil {
		return makeErr
	}

	return index.Make(kd)
}
