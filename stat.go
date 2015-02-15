package main

import "fmt"
import "github.com/kklipsch/cli"

func Stat(c *cli.Context, kayDir KayDir, index *Index) error {
	all, err := kayDir.ContentFiles()
	if err != nil {
		return err
	}

	PrintUnknownFiles(all, index)
	return nil
}

func PrintUnknownFiles(allFiles []File, index *Index) {
	for _, file := range allFiles {
		if !index.ContainsFile(file) {
			fmt.Printf("? %v\n", file)
		}
	}
}
