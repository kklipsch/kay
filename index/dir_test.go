package index

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
)

func TestCantCreateAnIndexOffANonExistentDir(t *testing.T) {
	if _, err := IndexDirectory(kaydir.KayDir("non-existent-path")); err == nil {
		t.Error("No error on non-existent path")
	}
}

func TestCantCreateAnIndexOffAFile(t *testing.T) {
	tempdir.In("file-fail-index", func(dir string) {
		file := path.Join(dir, "failfile")
		ioutil.WriteFile(file, []byte("failfile"), 600)

		if _, err := IndexDirectory(kaydir.KayDir(file)); err == nil {
			t.Errorf("No error on file path: %v", err)
		}
	})
}

func TestDirBasedConstructorPasses(t *testing.T) {
	tempdir.In("index-constructor", func(dir string) {
		Make(kaydir.KayDir(dir))
		if _, err := IndexDirectory(kaydir.KayDir(dir)); err != nil {
			t.Errorf("Error on construction: %v", err)
		}
	})
}

func TestAddWritesFile(t *testing.T) {
	withTempIndex("add-writes-file", func(index dirBasedIndex) {
		chap := chapter.Chapter("boo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notes")))
		if _, err := os.Stat(index.FullPath(chap)); err != nil {
			t.Errorf("Add Record did not write file: %v", err)
		}
	})
}

func TestAddChapter(t *testing.T) {
	withTempIndex("add-chapter", func(index dirBasedIndex) {
		chap := chapter.Chapter("foo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if !index.ContainsChapter(chap) {
			t.Errorf("Add record does not make the index contain the record")
		}
	})
}

func TestAddChapterUpdatesLastWritten(t *testing.T) {
	withTempIndex("last-written", func(index dirBasedIndex) {
		record := NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))
		before := time.Now()
		record, _ = index.AddChapter(chapter.Chapter("foo"), record)
		if !before.Before(record.LastWritten) {
			t.Errorf("index did not update last written")
		}
	})
}

func TestCannotAddSameChapterTwice(t *testing.T) {
	withTempIndex("last-written", func(index dirBasedIndex) {
		chap := chapter.Chapter("foo")
		index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if _, err := index.AddChapter(chap, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))); err == nil {
			t.Errorf("Allowed to add chapter twice.")
		}
	})
}

func TestIndexSurvivesMemoryLifetime(t *testing.T) {
	tempdir.In("index-survives", func(dir string) {
		kd := kaydir.KayDir(dir)
		chap := chapter.Chapter("foo")
		Make(kd)
		index1, _ := IndexDirectory(kd)
		if index2, _ := IndexDirectory(kd); index2 != nil {
			index2.AddChapter(chap, NewRecord(Year(1928), Note("notey")))
		}

		if !index1.ContainsChapter(chap) {
			t.Error("Index didn't contain expected file.")
		}
	})
}

func withTempIndex(label string, test func(dirBasedIndex)) error {
	return tempdir.In(label, func(dir string) {
		Make(kaydir.KayDir(dir))
		index, _ := IndexDirectory(kaydir.KayDir(dir))
		test(index.(dirBasedIndex))
	})
}
