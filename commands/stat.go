package commands

import "fmt"
import (
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func Stat(arguments Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	i, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	missing, err := GetMissingChapters(working, i)
	if err != nil {
		return err
	}

	for _, chap := range missing {
		fmt.Printf("? %v\n", chap)
	}

	return nil
}
