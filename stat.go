package main

import "fmt"
import (
	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/index"
)

func Stat(c *cli.Context, kayDir KayDir, i index.Index) error {
	all, err := kayDir.ContentFiles()
	if err != nil {
		return err
	}

	PrintUnknownFiles(all, i)
	return nil
}

func PrintUnknownFiles(allFiles []index.File, i index.Index) {
	for _, file := range allFiles {
		if !i.ContainsFile(file) {
			fmt.Printf("? %v\n", file)
		}
	}
}
