package chapter

import (
	"io/ioutil"
	"os"

	"github.com/kklipsch/kay/wd"
)

type Chapter string

func GetChaptersFromFileList(content []os.FileInfo) []Chapter {
	//is it bad to over allocate an array in go?
	chapters := make([]Chapter, 0)
	for _, f := range content {
		if !f.IsDir() {
			chapters = append(chapters, Chapter(f.Name()))
		}
	}
	return chapters
}

func GetChaptersFromPath(working wd.WorkingDirectory) ([]Chapter, error) {
	dir, err := ioutil.ReadDir(string(working))
	if err != nil {
		return nil, err
	}

	return GetChaptersFromFileList(dir), nil
}

func MapChaptersToString(chapters []Chapter) []string {
	strings := []string{}
	for _, c := range chapters {
		strings = append(strings, string(c))
	}
	return strings
}