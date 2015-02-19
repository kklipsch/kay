package main

import (
	"testing"
	"time"
)

func TestEmptyIndex(t *testing.T) {
	Assert(t, !EmptyIndex().ContainsFile(File("foo")), "Empty index shouldn't contain a file")
}

func TestCreateRecordAddsItToIndex(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")

	i.CreateRecord(Year(1952), f, Note("foo"), Tag("1"), Tag("b"))
	Assert(t, i.ContainsFile(f), "New record should add to index")
}

func TestCreateRecordUpdatesRecordCorrectly(t *testing.T) {
	f := File("foo")
	before := time.Now()

	r, _ := EmptyIndex().CreateRecord(Year(1952), f, Note("foo"), Tag("1"), Tag("b"))

	Assert(t, r.Note == "foo", "Note set")
	Assert(t, r.DateAdded.After(before), "Date Added is set")
	Assert(t, r.Year == Year(1952), "Year is set")
	Assert(t, r.File() == f, "File is set")
	Assert(t, r.HasTag("1"), "Tag set 1")
	Assert(t, r.HasTag("b"), "Tag set b")
	Assert(t, !r.HasTag("moo"), "Moo is there")
}

func TestCreateRecordForSameFileTwice(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")

	i.CreateRecord(Year(1952), f, "")
	_, err := i.CreateRecord(Year(1953), f, "")
	FailIfNoError(t, err, "Should not be able to add file twice")
}

func TestGetRecordForFileFile(t *testing.T) {
	i := EmptyIndex()
	f1 := File("foo1")
	f2 := File("foo2")

	i.CreateRecord(Year(1952), f1, "")
	i.CreateRecord(Year(1962), f2, "")

	Assert(t, i.Get("foo1").File() == f1, "First file get")
	Assert(t, i.Get("foo2").File() == f2, "Second file get")
	Assert(t, i.Get("foo3") == nil, "Get nonexistent file")
}

func TestGottenFileChangesAreRepresented(t *testing.T) {
	i := EmptyIndex()
	f := File("foo1")
	i.CreateRecord(Year(1952), f, "")

	r1 := i.Get(f)
	r2 := i.Get(f)

	Assert(t, r1 == r2, "Pointer equality")
	Assert(t, r1.Year == r2.Year, "Same Values")

	r1.Year = Year(1953)
	Assert(t, r1.Year == r2.Year, "Same Values after swap")
}

func TestMoveFile(t *testing.T) {
	i := EmptyIndex()
	r, _ := i.CreateRecord(Year(1952), File("foo"), "")
	i.MoveRecord(r, File("moo"))

	Assert(t, r.File() == File("moo"), "File was changed")
	Assert(t, i.ContainsFile("moo"), "New file is found")
	Assert(t, !i.ContainsFile("foo"), "File was rightly gone after move")
}

func TestMoveNonExistentFile(t *testing.T) {
	i := EmptyIndex()

	r, _ := i.CreateRecord(Year(1952), File("foo"), "")
	i.DeleteRecord(r)
	_, err := i.MoveRecord(r, File("foo2"))

	FailIfNoError(t, err, "Shouldn't be able to non existing record.")
}

func TestMoveIntoExistingFile(t *testing.T) {
	i := EmptyIndex()
	r, _ := i.CreateRecord(Year(1952), File("foo"), "")
	i.CreateRecord(Year(1953), File("foo2"), "")

	_, err := i.MoveRecord(r, File("foo2"))
	FailIfNoError(t, err, "Shouldn't be able to move into an existing file")
}

func TestSetRecord(t *testing.T) {
	r, _ := EmptyIndex().CreateRecord(Year(1952), File("foo"), "")

	i := EmptyIndex()
	i.SetRecord(r, File("moo"))

	Assert(t, r.File() == File("moo"), "File was changed")
	Assert(t, i.ContainsFile("moo"), "New file is found")
	Assert(t, !i.ContainsFile("foo"), "File still not in index")
}

func TestSetRecordToExistingFile(t *testing.T) {
	r, _ := EmptyIndex().CreateRecord(Year(1952), File("foo"), "")

	i := EmptyIndex()
	i.CreateRecord(Year(1952), File("moo"), "")
	_, err := i.SetRecord(r, File("moo"))
	FailIfNoError(t, err, "Should not be allowed to set a record to an existing file.")
}

func TestDeleteFile(t *testing.T) {
	i := EmptyIndex()
	f := File("foo1")
	r, _ := i.CreateRecord(Year(1952), f, "")
	i.DeleteRecord(r)

	Assert(t, !i.ContainsFile(f), "File removed")
}

func TestCantDeleteNonExistentFile(t *testing.T) {
	r, _ := EmptyIndex().CreateRecord(Year(1952), File("foo"), "")
	FailIfNoError(t, EmptyIndex().DeleteRecord(r), "Should not be able to delete non existent record.")
}

func TestHasTags(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	r, _ := i.CreateRecord(Year(1952), f, "")

	r.Tags = append(r.Tags, Tag("foo"))
	r.Tags = append(r.Tags, Tag("bar"))

	AssertTags(t, r.Tags, []Tag{Tag("foo"), Tag("bar")}, "add tag works")
	Assert(t, r.HasTag("bar"), "Has tag works")

	r.Tags = []Tag{}
	AssertTags(t, r.Tags, []Tag{}, "Clear tags empties them")
	Assert(t, !r.HasTag("bar"), "Clear really clears")
}

func TestUpdateNote(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	r, _ := i.CreateRecord(Year(1952), f, "note 1")

	r.Note = Note("note 2")
	Assert(t, r.Note == Note("note 2"), "note is updated")
}

func TestSwapYear(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	r, _ := i.CreateRecord(Year(1952), f, "")

	r.Year = Year(1953)
	Assert(t, r.Year == Year(1953), "year is updated")
}
