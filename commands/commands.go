package commands

import (
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

type Arguments interface{}

func RunCommand(arguments Arguments, command func(Arguments, kaydir.KayDir, wd.WorkingDirectory) error) error {
	working, wdErr := wd.Get()
	if wdErr != nil {
		return wdErr
	}

	kd, kdErr := kaydir.Get(working)
	if kdErr != nil {
		return kdErr
	}

	return command(arguments, kd, working)
}
