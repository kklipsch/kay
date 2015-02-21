package main

import (
	"fmt"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/index"
)

func Add(c *cli.Context, kayDir KayDir, i index.Index) error {
	file := index.File(c.Args().First())
	year, year_parsed, err := ParseYear(file)
	if err != nil {
		return err
	}

	if !year_parsed {
		return fmt.Errorf("No year provided or parsed!")
	}

	_, addErr := i.AddRecord(file, index.NewRecord(year, index.Note("")))
	return addErr
}

//stubbed out
func ParseYear(file index.File) (index.Year, bool, error) {
	return 1941, true, nil
}
