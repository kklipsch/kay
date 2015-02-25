package commands

import "fmt"
import (
	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
)

func Stat(c *cli.Context, kd kaydir.KayDir, i index.Index) error {
	all, err := chapter.GetChaptersFromPath(".")
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
