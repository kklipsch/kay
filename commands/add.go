package commands

import (
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func Add(arguments Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	i, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	chapter := chapter.Chapter("1941.test1.2015_02.doc")
	_, addErr := i.AddChapter(chapter, index.NewRecord(1941, index.Note("")))
	return addErr
}

//stubbed out
func ParseYear(chapter chapter.Chapter) (index.Year, bool, error) {
	return 1941, true, nil
}
