package index

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/kklipsch/kay/tempdir"
)

func TestCantCreateAnIndexOffANonExistentDir(t *testing.T) {
	if _, err := IndexDirectory("non-existent-path"); err == nil {
		t.Error("No error on non-existent path")
	}
}

func TestCantCreateAnIndexOffAFile(t *testing.T) {
	tempdir.In("file-fail-index", func(dir string) {
		file := path.Join(dir, "failfile")
		ioutil.WriteFile(file, []byte("failfile"), 600)

		if _, err := IndexDirectory(file); err == nil {
			t.Errorf("No error on file path: %v", err)
		}
	})
}

func TestDirBasedConstructorPasses(t *testing.T) {
	tempdir.In("index-constructor", func(dir string) {
		if _, err := IndexDirectory(dir); err != nil {
			t.Errorf("Error on construction: %v", err)
		}
	})
}

func TestAddWritesFile(t *testing.T) {
	withTempIndex("add-writes-file", func(index dirBasedIndex) {
		file := File("boo")
		index.AddRecord(file, NewRecord(Year(1928), Note("notes")))
		if _, err := os.Stat(index.FullPath(file)); err != nil {
			t.Errorf("Add Record did not write file: %v", err)
		}
	})
}

func TestAddRecord(t *testing.T) {
	withTempIndex("add-record", func(index dirBasedIndex) {
		file := File("foo")
		index.AddRecord(file, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if !index.ContainsFile(file) {
			t.Errorf("Add record does not make the index contain the record")
		}
	})
}

func TestAddRecordUpdatesLastWritten(t *testing.T) {
	withTempIndex("last-written", func(index dirBasedIndex) {
		record := NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))
		before := time.Now()
		record, _ = index.AddRecord(File("foo"), record)
		if !before.Before(record.LastWritten) {
			t.Errorf("index did not update last written")
		}
	})
}

func TestCannotAddSameFileTwice(t *testing.T) {
	withTempIndex("last-written", func(index dirBasedIndex) {
		file := File("foo")
		index.AddRecord(file, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2")))
		if _, err := index.AddRecord(file, NewRecord(Year(1928), Note("notey"), Tag("tag1"), Tag("tag2"))); err == nil {
			t.Errorf("Allowed to add file twice.")
		}
	})
}

func TestIndexSurvivesMemoryLifetime(t *testing.T) {
	tempdir.In("index-survives", func(dir string) {
		file := File("foo")
		index1, _ := IndexDirectory(dir)
		if index2, _ := IndexDirectory(dir); index2 != nil {
			index2.AddRecord(file, NewRecord(Year(1928), Note("notey")))
		}

		if !index1.ContainsFile(file) {
			t.Error("Index didn't contain expected file.")
		}
	})
}

func withTempIndex(label string, test func(dirBasedIndex)) error {
	return tempdir.In(label, func(dir string) {
		index, _ := IndexDirectory(dir)
		test(index.(dirBasedIndex))
	})
}
