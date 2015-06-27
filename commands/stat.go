package commands

import "fmt"
import (
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

//Stat prints any chapters that exist in the working directory but are not indexed.
func Stat(kd kaydir.KayDir, working wd.WorkingDirectory) error {
	i, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	missing, err := getMissingChapters(working, i)
	if err != nil {
		return err
	}

	for _, chap := range missing {
		fmt.Printf("? %v\n", chap)
	}

	return nil
}
