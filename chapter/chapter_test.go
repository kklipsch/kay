package chapter

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/kklipsch/kay/tempdir"
)

func TestDeepEqual(t *testing.T) {
	start := []string{"test1", "test2"}

	second := []string{}
	for _, c := range start {
		second = append(second, string(c))
	}

	if !reflect.DeepEqual(start, second) {
		t.Errorf("Deep equal is useless %v:%v", start, second)
	}
}

func TestGetChaptersFromPath(t *testing.T) {
	tempdir.In("get-chapters-from-path", func(dir string) {
		ioutil.WriteFile(filepath.Join(dir, "test1"), []byte("test1"), 0777)
		ioutil.WriteFile(filepath.Join(dir, "test2"), []byte("test2"), 0777)

		expected := []string{"test1", "test2"}
		chapters, _ := GetChaptersFromPath(dir)
		if !reflect.DeepEqual(MapChaptersToString(chapters), expected) {
			t.Errorf("Expected %v Got %v", expected, chapters)
		}
	})
}

func TestGetChaptersFromPathFiltersDirectories(t *testing.T) {
	tempdir.In("get-chapters-filters-dirs", func(dir string) {
		ioutil.WriteFile(filepath.Join(dir, "test1"), []byte("test1"), 0777)
		os.MkdirAll(filepath.Join(dir, "test2"), 0755)
		ioutil.WriteFile(filepath.Join(dir, "test3"), []byte("test3"), 0777)

		expected := []string{"test1", "test3"}
		chapters, _ := GetChaptersFromPath(dir)
		if !reflect.DeepEqual(MapChaptersToString(chapters), expected) {
			t.Errorf("Expected %v Got %v", expected, chapters)
		}
	})
}
