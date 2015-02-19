package main

import (
	"io/ioutil"
	"os"
)

func GetFilesFromDir(path string) ([]File, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return getFilesFromList(path, dir), nil
}

func getFilesFromList(path string, content []os.FileInfo) []File {
	files := []File{}
	for _, f := range content {
		if !f.IsDir() {
			files = append(files, File(f.Name()))
		}
	}
	return files
}
