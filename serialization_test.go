package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kklipsch/kay/tempdir"
)

func TestSerializingOneRecordBackAndForth(t *testing.T) {
	index := EmptyIndex()
	record, err := index.CreateRecord(Year(1942), File("file1"), Note(""), Tag("Foo"), Tag("bar"))
	FailIfError(t, err, "Could not create record for serialize one test")

	before := time.Now()

	ser, serErr := SerializeRecord(record)
	FailIfError(t, serErr, "Could not serialize one record")
	sut, deserErr := DeserializeRecord(File("file1"), EmptyIndex(), ser)
	FailIfError(t, deserErr, "Could not deserialize one record")
	AssertRecords(t, sut, record, "Serialize one record failed")
	Assert(t, record.LastWritten.After(before), "Last Written is updated")
}

func TestSerializationBackAndForth(t *testing.T) {
	tempdir.In("serialization_back_forth", func(dir string) {
		indexdir := filepath.Join(dir, "foo", "bar")
		FailIfError(t, os.MkdirAll(dir, 0755), "couldnt create index dir for serialization")

		file1Name := File("file1")
		file2Name := File("file2")

		index := EmptyIndex()
		file1, file1Err := index.CreateRecord(Year(1942), file1Name, Note(""))
		FailIfError(t, file1Err, "Cant create first serialization test record")
		file2, file2Err := index.CreateRecord(Year(1943), file2Name, Note(""))
		FailIfError(t, file2Err, "Cant create second serialization test record")

		FailIfError(t, WriteIndex(indexdir, index), "Could not serialize index")
		fileIndex, serErr := BuildIndex(indexdir)
		FailIfError(t, serErr, "Could not deserialize index")

		serFile1 := fileIndex.Get(file1Name)
		serFile2 := fileIndex.Get(file2Name)

		if serFile1 == nil || serFile2 == nil {
			t.Fatalf("Unable to get file1 (%v) or file2 (%v)", serFile1, serFile2)
		}

		AssertRecords(t, file1, serFile1, "serialization file 1 match")
		AssertRecords(t, file2, serFile2, "serialization file 1 match")
	})
}
