package commands

import "fmt"
import (
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func Stat(arguments Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	i, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	all, err := chapter.GetChaptersFromPath(working)
	if err != nil {
		return err
	}

	PrintUnknownChapters(all, i)
	return nil
}

func PrintUnknownChapters(allChapters []chapter.Chapter, i index.Index) {
	for _, chap := range allChapters {
		if !i.ContainsChapter(chap) {
			fmt.Printf("? %v\n", chap)
		}
	}
}
