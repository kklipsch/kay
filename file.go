package main

import (
	"io/ioutil"
	"os"

	"github.com/kklipsch/kay/index"
)

func GetFilesFromDir(path string) ([]index.File, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return getFilesFromList(path, dir), nil
}

func getFilesFromList(path string, content []os.FileInfo) []index.File {
	//is it bad to over allocate an array in go?
	files := make([]index.File, 0)
	for _, f := range content {
		if !f.IsDir() {
			files = append(files, index.File(f.Name()))
		}
	}
	return files
}
