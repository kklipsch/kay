package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilesFromDir(t *testing.T) {
	InTempDir(t, "get files from dir", func(dir string) {
		ioutil.WriteFile(filepath.Join(dir, "test1"), []byte("test1"), 0777)
		os.MkdirAll(filepath.Join(dir, "test2"), 0755) //dirs should be filtered
		ioutil.WriteFile(filepath.Join(dir, "test3"), []byte("test3"), 0777)

		files, _ := GetFilesFromDir(dir)
		assertFiles(t, files, []string{"test1", "test3"}, dir, "Get only files from directory")

	})
}

func assertFiles(t *testing.T, files []File, expected []string, dir string, msg string) {
	if len(files) != len(expected) {
		t.Fatalf("Expected %v Got %v: %s", expected, files, msg)
	}

	for i, v := range files {
		fullpath := expected[i]
		if fullpath != string(v) {
			t.Fatalf("Expected %v Got %v: %s", fullpath, v, msg)
		}
	}
}
