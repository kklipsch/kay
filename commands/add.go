package commands

import (
	"fmt"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
)

func Add(c *cli.Context, kd kaydir.KayDir, i index.Index) error {
	chapter := chapter.Chapter(c.Args().First())
	year, year_parsed, err := ParseYear(chapter)
	if err != nil {
		return err
	}

	if !year_parsed {
		return fmt.Errorf("No year provided or parsed!")
	}

	_, addErr := i.AddChapter(chapter, index.NewRecord(year, index.Note("")))
	return addErr
}

//stubbed out
func ParseYear(chapter chapter.Chapter) (index.Year, bool, error) {
	return 1941, true, nil
}
