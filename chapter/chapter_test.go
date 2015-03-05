package chapter

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
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
	tempdir.TempWd(func(dir wd.WorkingDirectory) {
		ioutil.WriteFile(dir.Path("test1"), []byte("test1"), 0777)
		ioutil.WriteFile(dir.Path("test2"), []byte("test2"), 0777)

		expected := []string{"test1", "test2"}
		chapters, _ := GetChaptersFromPath(dir)
		if !reflect.DeepEqual(MapChaptersToString(chapters), expected) {
			t.Errorf("Expected %v Got %v", expected, chapters)
		}
	})
}

func TestGetChaptersFromPathFiltersDirectories(t *testing.T) {
	tempdir.TempWd(func(dir wd.WorkingDirectory) {
		ioutil.WriteFile(dir.Path("test1"), []byte("test1"), 0777)
		os.MkdirAll(dir.Path("test2"), 0755)
		ioutil.WriteFile(dir.Path("test3"), []byte("test3"), 0777)

		expected := []string{"test1", "test3"}
		chapters, _ := GetChaptersFromPath(dir)
		if !reflect.DeepEqual(MapChaptersToString(chapters), expected) {
			t.Errorf("Expected %v Got %v", expected, chapters)
		}
	})
}
