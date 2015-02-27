package commands

import (
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

type Arguments struct {
	Chapters []chapter.Chapter
	Year     index.Year
}

func RunWithKayDir(arguments Arguments, working wd.WorkingDirectory, command func(Arguments, kaydir.KayDir, wd.WorkingDirectory) error) error {
	kd, kdErr := kaydir.Get(working)
	if kdErr != nil {
		return kdErr
	}

	return command(arguments, kd, working)
}
