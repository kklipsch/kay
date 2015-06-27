package chapter

import (
	"io/ioutil"
	"os"

	"github.com/kklipsch/kay/wd"
)

//Chapter a chapter represents a single version of a given year's story.
type Chapter string

func filesToChapters(content []os.FileInfo) []Chapter {
	var chapters []Chapter
	for _, f := range content {
		if !f.IsDir() {
			chapters = append(chapters, Chapter(f.Name()))
		}
	}
	return chapters
}

//GetChaptersFromPath will get every file in the working directory as a chapter.
func GetChaptersFromPath(working wd.WorkingDirectory) ([]Chapter, error) {
	dir, err := ioutil.ReadDir(string(working))
	if err != nil {
		return nil, err
	}

	return filesToChapters(dir), nil
}

func chaptersAsStrings(chapters []Chapter) []string {
	strings := []string{}
	for _, chaptersAsStrings := range chapters {
		strings = append(strings, string(chaptersAsStrings))
	}
	return strings
}
