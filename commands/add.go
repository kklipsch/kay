package commands

import (
	"fmt"
	"regexp"
	"strconv"

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

	toAdd, addErr := getChaptersToAdd(arguments, working, i)
	if addErr != nil {
		return addErr
	}

	return CompositeError(addChapters(i, toAdd, getYear(arguments), getTags(arguments), getNotes(arguments)))
}

func getChaptersToAdd(arguments Arguments, working wd.WorkingDirectory, i index.Index) ([]chapter.Chapter, error) {
	if len(arguments.Chapters) > 0 {
		return arguments.Chapters, nil
	} else {
		return GetMissingChapters(working, i)
	}
}

func addChapters(i index.Index, chapters []chapter.Chapter, year yearChoice, tags tagChoice, notes notesChoice) []error {

	errors := make([]error, 0)
	for _, chapter := range chapters {
		if addErr := addChapter(i, chapter, year, tags, notes); addErr != nil {
			errors = append(errors, addErr)
		}
	}

	return errors
}

func addChapter(i index.Index, chap chapter.Chapter, year yearChoice, tags tagChoice, notes notesChoice) error {

	y, yerr := year(chap)
	if yerr != nil {
		return yerr
	}

	t, terr := tags(chap)
	if terr != nil {
		return terr
	}

	n, nerr := notes(chap)
	if nerr != nil {
		return nerr
	}

	if _, addErr := i.AddChapter(chap, index.NewRecord(y, n, t...)); addErr != nil {
		return addErr
	}

	return nil
}

type yearChoice func(chapter.Chapter) (index.Year, error)
type tagChoice func(chapter.Chapter) ([]index.Tag, error)
type notesChoice func(chapter.Chapter) (index.Note, error)

func getYear(arguments Arguments) yearChoice {
	return func(chap chapter.Chapter) (index.Year, error) {
		if arguments.Year != index.EmptyYear {
			return arguments.Year, nil
		} else {
			return parseYear(chap)
		}
	}
}

func getTags(arguments Arguments) tagChoice {
	return func(chap chapter.Chapter) ([]index.Tag, error) {
		return []index.Tag{}, nil
	}
}

func getNotes(arguments Arguments) notesChoice {
	return func(chap chapter.Chapter) (index.Note, error) {
		return index.Note(""), nil
	}
}

func parseYear(chap chapter.Chapter) (index.Year, error) {
	yearReg := regexp.MustCompile(`^([0-9]{4})\..*`)

	yearString := yearReg.FindStringSubmatch(string(chap))
	if yearString == nil {
		return index.EmptyYear, fmt.Errorf("Could not find year in: %v", chap)
	}

	yearNum, convErr := strconv.Atoi(yearString[1])
	if convErr != nil {
		return index.EmptyYear, convErr
	}

	return index.Year(yearNum), nil
}
