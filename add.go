package main

import (
	"fmt"

	"github.com/kklipsch/cli"
)

func Add(c *cli.Context, kayDir KayDir, index *Index) error {
	file := File(c.Args().First())
	year, year_parsed, err := ParseYear(file)
	if err != nil {
		return err
	}

	if !year_parsed {
		return fmt.Errorf("No year provided or parsed!")
	}

	index.CreateRecord(year, file, Note(""))
	return nil
}

//stubbed out
func ParseYear(file File) (Year, bool, error) {
	return 1941, true, nil
}
