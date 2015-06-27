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

//AddArguments are anything necessary to add chapters
type AddArguments struct {
	Chapters []chapter.Chapter
	Year     index.Year
	Tags     []index.Tag
	Note     index.Note
}

//Add will add any missing chapters with the supplied arguments
func Add(arguments AddArguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	i, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	toAdd, addErr := getChaptersToAdd(arguments, working, i)
	if addErr != nil {
		return addErr
	}

	return compositeError(addChapters(i, toAdd, getYear(arguments), getTags(arguments), getNotes(arguments)))
}

func getChaptersToAdd(arguments AddArguments, working wd.WorkingDirectory, i index.Index) ([]chapter.Chapter, error) {
	if len(arguments.Chapters) > 0 {
		return arguments.Chapters, nil
	}

	return getMissingChapters(working, i)
}

func addChapters(i index.Index, chapters []chapter.Chapter, year yearChoice, tags tagChoice, notes notesChoice) []error {

	var errors []error
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

func getYear(arguments AddArguments) yearChoice {
	return func(chap chapter.Chapter) (index.Year, error) {
		if arguments.Year != index.EmptyYear {
			return arguments.Year, nil
		}

		return parseYear(chap)
	}
}

func getTags(arguments AddArguments) tagChoice {
	return func(chap chapter.Chapter) ([]index.Tag, error) {
		if len(arguments.Tags) > 0 {
			return arguments.Tags, nil
		}

		tag, err := parseTag(chap)
		if err != nil {
			return []index.Tag{}, err
		}

		return []index.Tag{tag}, nil
	}
}

func getNotes(arguments AddArguments) notesChoice {
	return func(chap chapter.Chapter) (index.Note, error) {
		return arguments.Note, nil
	}
}

func parseTag(chap chapter.Chapter) (index.Tag, error) {
	tagRexp := regexp.MustCompile(`^[0-9]{4}\.(.*)\..*$`)
	tagString := tagRexp.FindStringSubmatch(string(chap))
	if tagString == nil {
		return index.Tag(""), nil
	}

	return index.Tag(tagString[1]), nil
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
