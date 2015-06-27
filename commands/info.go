package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func parseMode(mode string) (func(chapter.Chapter, *index.Record) (string, error), error) {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "normal", "":
		return normal, nil
	case "json":
		return jsonOut, nil
	case "year":
		return year, nil
	case "tags":
		return tags, nil
	case "note":
		return notes, nil
	case "added":
		return added, nil
	case "written":
		return written, nil
	}

	return nil, fmt.Errorf("Unknown mode %s", mode)
}

func normal(chapter chapter.Chapter, record *index.Record) (string, error) {
	return fmt.Sprintf("Name:%s Year:%v Note:%s Tags:%s Added:%s Last Updated:%s",
		chapter,
		record.Year,
		record.Note,
		record.Tags,
		localTime(record.DateAdded),
		localTime(record.LastWritten)), nil
}

func jsonOut(chapter chapter.Chapter, record *index.Record) (string, error) {
	json, jsonErr := json.Marshal(&record)
	if jsonErr != nil {
		return "", jsonErr
	}

	return string(json), nil
}

func year(chapter chapter.Chapter, record *index.Record) (string, error) {
	return fmt.Sprintf("%v", record.Year), nil
}

func tags(chapter chapter.Chapter, record *index.Record) (string, error) {
	tags := []string{}
	for _, tag := range record.Tags {
		tags = append(tags, fmt.Sprintf("%v", tag))
	}

	return strings.Join(tags, ","), nil
}

func notes(chapter chapter.Chapter, record *index.Record) (string, error) {
	return fmt.Sprintf("%v", record.Note), nil
}

func added(chapter chapter.Chapter, record *index.Record) (string, error) {
	return localTime(record.DateAdded), nil
}

func written(chapter chapter.Chapter, record *index.Record) (string, error) {
	return localTime(record.LastWritten), nil
}

func localTime(t time.Time) string {
	local := t.Local()
	return local.Format("Mon, 01/02/06, 03:04PM")
}

//Info prints the index informaiton for the given chapters, output is based on the mode.
func Info(chapters []chapter.Chapter, mode string, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	display, parseErr := parseMode(mode)
	if parseErr != nil {
		return parseErr
	}

	index, indexErr := index.Get(kd)
	if indexErr != nil {
		return indexErr
	}

	if len(chapters) == 0 {
		chapters = index.AllIndexed()
	}

	errors := []error{}
	for _, chapter := range chapters {
		if err := infoChapter(chapter, index, display); err != nil {
			errors = append(errors, err)
		}
	}

	return compositeError(errors)
}

func infoChapter(chapter chapter.Chapter, index index.Index, display func(chapter.Chapter, *index.Record) (string, error)) error {
	if !index.ContainsChapter(chapter) {
		return fmt.Errorf("%v is not indexed", chapter)
	}

	record, getErr := index.GetRecord(chapter)
	if getErr != nil {
		return getErr
	}

	output, displayErr := display(chapter, record)
	if displayErr != nil {
		return displayErr
	}

	fmt.Println(output)
	return nil
}
