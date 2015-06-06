package commands

import (
	"fmt"
	"strings"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kayignore"
	"github.com/kklipsch/kay/wd"
)

func CompositeError(errors []error) error {
	if len(errors) > 0 {
		var messages []string
		for _, err := range errors {
			messages = append(messages, err.Error())
		}
		return fmt.Errorf("%v", strings.Join(messages, "\n"))
	} else {
		return nil
	}
}

func GetMissingChapters(working wd.WorkingDirectory, i index.Index) ([]chapter.Chapter, error) {
	var chapters []chapter.Chapter

	ki, err := kayignore.Get(working)
	if err != nil {
		return chapters, err
	}

	all, err := chapter.GetChaptersFromPath(working)
	if err != nil {
		return chapters, err
	}

	for _, chap := range all {
		if !i.ContainsChapter(chap) && !ki.Ignore(chap) {
			chapters = append(chapters, chap)
		}
	}

	return chapters, nil
}
