package main

import "fmt"
import (
	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
)

func Stat(c *cli.Context, kayDir KayDir, i index.Index) error {
	all, err := kayDir.ContentChapters()
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
