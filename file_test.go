package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	InTempDir(t, "file exists", func(dir string) {
		exists, err := FileExists(dir)
		FailIfError(t, err, "Temp dir exists error")
		Assert(t, exists, "Temp dir does not exist")

		exists, err = FileExists(filepath.Join(dir, "foo"))
		FailIfError(t, err, "Non existent file exists err")
		Assert(t, !exists, "Non existent file exists")
	})
}

func TestGetFilesFromDir(t *testing.T) {
	InTempDir(t, "get files from dir", func(dir string) {
		FailIfError(t, ioutil.WriteFile(filepath.Join(dir, "test1"), []byte("test1"), 0777), "couldnt create test1")
		FailIfError(t, os.MkdirAll(filepath.Join(dir, "test2"), 0755), "Couldnt create test2")
		FailIfError(t, ioutil.WriteFile(filepath.Join(dir, "test3"), []byte("test3"), 0777), "Couldnt make test3")

		files, err := GetFilesFromDir(dir)
		FailIfError(t, err, "Couldnt get files")
		assertFiles(t, files, []string{"test1", "test3"}, dir, "Get files")

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
