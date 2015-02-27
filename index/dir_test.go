package index

import (
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
	"time"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
)

func TestCantCreateAnIndexOffANonExistentDir(t *testing.T) {
	if _, err := Get(kaydir.KayDir("non-existent-path")); err == nil {
		t.Error("No error on non-existent path")
	}
}

func TestCantCreateAnIndexOffAFile(t *testing.T) {
	tempdir.In(func(dir string) {
		file := path.Join(dir, "failfile")
		ioutil.WriteFile(file, []byte("failfile"), 600)

		if _, err := Get(kaydir.KayDir(file)); err == nil {
			t.Errorf("No error on file path: %v", err)
		}
	})
}

func TestDirBasedConstructorPasses(t *testing.T) {
	tempdir.In(func(dir string) {
		Make(kaydir.KayDir(dir))
		if _, err := Get(kaydir.KayDir(dir)); err != nil {
			t.Errorf("Error on construction: %v", err)
		}
	})
}

func TestAddWritesFile(t *testing.T) {
	withTempIndex(func(index dirBasedIndex) {
		chap := chapter.Chapter("boo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notes")))
		if _, err := os.Stat(index.FullPath(chap)); err != nil {
			t.Errorf("Add Record did not write file: %v", err)
		}
	})
}

func TestAddChapter(t *testing.T) {
	withTempIndex(func(index dirBasedIndex) {
		chap := chapter.Chapter("foo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if !index.ContainsChapter(chap) {
			t.Errorf("Add record does not make the index contain the record")
		}
	})
}

func TestAddChapterUpdatesLastWritten(t *testing.T) {
	withTempIndex(func(index dirBasedIndex) {
		record := NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))
		before := time.Now()
		record, _ = index.AddChapter(chapter.Chapter("foo"), record)
		if !before.Before(record.LastWritten) {
			t.Errorf("index did not update last written")
		}
	})
}

func TestCannotAddSameChapterTwice(t *testing.T) {
	withTempIndex(func(index dirBasedIndex) {
		chap := chapter.Chapter("foo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if _, err := index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))); err == nil {
			t.Errorf("Allowed to add chapter twice.")
		}
	})
}

func TestIndexSurvivesMemoryLifetime(t *testing.T) {
	tempdir.In(func(dir string) {
		kd := kaydir.KayDir(dir)
		chap := chapter.Chapter("foo")
		Make(kd)
		index1, _ := Get(kd)
		if index2, _ := Get(kd); index2 != nil {
			index2.AddChapter(chap, NewRecord(Year(1928), Note("notey")))
		}

		if !index1.ContainsChapter(chap) {
			t.Error("Index didn't contain expected file.")
		}
	})
}

func TestGetRecord(t *testing.T) {
	tempdir.In(func(dir string) {
		kd := kaydir.KayDir(dir)
		chap := chapter.Chapter("foo")

		Make(kd)

		index1, _ := Get(kd)
		rec, _ := index1.AddChapter(chap, NewRecord(Year(1928), Note("notey")))

		index2, _ := Get(kd)
		got, getErr := index2.GetRecord(chap)

		if getErr != nil {
			t.Errorf("Error on get: %v", getErr)
		}

		if !reflect.DeepEqual(rec, got) {
			t.Errorf("Expected %v Got %v", rec, got)
		}
	})
}

func TestGetNonexistantRecord(t *testing.T) {
	withTempIndex(func(index dirBasedIndex) {
		rec, err := index.GetRecord(chapter.Chapter("foo"))

		if rec != nil {
			t.Errorf("Record should be empty: %v", rec)
		}

		if err == nil {
			t.Errorf("Error should not be empty")
		}

	})
}

func withTempIndex(test func(dirBasedIndex)) error {
	return tempdir.In(func(dir string) {
		Make(kaydir.KayDir(dir))
		index, _ := Get(kaydir.KayDir(dir))
		test(index.(dirBasedIndex))
	})
}
