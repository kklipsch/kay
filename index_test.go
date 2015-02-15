package main

import (
	"testing"
	"time"
)

func TestCreateRecord(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	before := time.Now()

	Assert(t, !i.ContainsFile(f), "Empty index shouldn't contain a file")

	r, err := i.CreateRecord(Year(1952), f, Note("foo"), Tag("1"), Tag("b"))
	FailIfError(t, err, "Create record failed")

	Assert(t, r.Note == "foo", "Note set")
	Assert(t, r.DateAdded.After(before), "Date Added is set")
	Assert(t, r.Year == Year(1952), "Year is set")
	Assert(t, r.File() == f, "File is set")
	Assert(t, r.HasTag("1"), "Tag set 1")
	Assert(t, r.HasTag("b"), "Tag set b")
	Assert(t, !r.HasTag("moo"), "Moo is there")

	Assert(t, i.ContainsFile(f), "New record should add to index")
}

func TestCreateSameFileTwice(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	Assert(t, !i.ContainsFile(f), "Empty index shouldn't contain a file")
	_, err := i.CreateRecord(Year(1952), f, "")
	Assert(t, i.ContainsFile(f), "New record should add to index")
	_, err = i.CreateRecord(Year(1953), f, "")
	FailIfNoError(t, err, "Should not be able to add file twice")
}

func TestGetFile(t *testing.T) {
	i := EmptyIndex()
	f1 := File("foo1")
	f2 := File("foo2")

	_, err := i.CreateRecord(Year(1952), f1, "")
	FailIfError(t, err, "Could not create record for getting")
	_, err = i.CreateRecord(Year(1962), f2, "")
	FailIfError(t, err, "Could not create second record for getting")

	Assert(t, i.Get("foo1").File() == f1, "First file get")
	Assert(t, i.Get("foo2").File() == f2, "Second file get")
	Assert(t, i.Get("foo3") == nil, "Get nonexistent file")
}

func TestGottenFileChangesAreRepresented(t *testing.T) {
	i := EmptyIndex()
	f := File("foo1")
	_, err := i.CreateRecord(Year(1952), f, "")
	FailIfError(t, err, "Failed ot create error for change test")

	r1 := i.Get(f)
	r2 := i.Get(f)

	Assert(t, r1 == r2, "Pointer equality")
	Assert(t, r1.Year == r2.Year, "Same Values")
	r1.Year = Year(1953)
	Assert(t, r1.Year == r2.Year, "Same Values after swap")
}

func TestMoveFile(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Unable to create record for move file test")

	Assert(t, r.File() == File("foo"), "File has original name")
	Assert(t, i.ContainsFile("foo"), "Found file before move")
	_, err = i.MoveRecord(r, File("moo"))
	FailIfError(t, err, "Failed to move file")

	Assert(t, r.File() == File("moo"), "File was changed")
	Assert(t, i.ContainsFile("moo"), "New file is found")
	Assert(t, !i.ContainsFile("foo"), "File was rightly gone after move")
}

func TestMoveNonExistentFile(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Failed on create record for move non-existent file")

	err = i.DeleteRecord(r)
	FailIfError(t, err, "Failed to delete record for move non-existent file")

	Assert(t, !i.ContainsFile(File("foo")), "File does not exist before move.")
	_, err = i.MoveRecord(r, File("foo2"))
	FailIfNoError(t, err, "Shouldn't be able to move record into a non existing file")
}

func TestMoveIntoExistingFile(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Unable to create first file for move existing file test")

	_, err = i.CreateRecord(Year(1953), File("foo2"), "")
	FailIfError(t, err, "Unable to create second file for move existing file test")

	_, err = i.MoveRecord(r, File("foo2"))
	FailIfNoError(t, err, "Shouldn't be able to move into an existing file")
}

func TestSetRecord(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Unable to create record for set file test")
	FailIfError(t, i.DeleteRecord(r), "Unable to delete record in set record test")

	Assert(t, !i.ContainsFile("foo"), "Found file before move")
	_, err = i.SetRecord(r, File("moo"))
	FailIfError(t, err, "Failed to set file")

	Assert(t, r.File() == File("moo"), "File was changed")
	Assert(t, i.ContainsFile("moo"), "New file is found")
	Assert(t, !i.ContainsFile("foo"), "File still not in index")
}

func TestSetRecordToExistingFile(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Unable to create record for set file test")
	FailIfError(t, i.DeleteRecord(r), "Unable to delete record in set record test")

	_, err = i.CreateRecord(Year(1952), File("moo"), "")
	FailIfError(t, err, "Unable to create record 2 for set file test")
	_, err = i.SetRecord(r, File("moo"))
	FailIfNoError(t, err, "Allowed to set record to existing")
}

func TestDeleteFile(t *testing.T) {
	i := EmptyIndex()
	_, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Failed on first create for delete file test")

	r, err2 := i.CreateRecord(Year(1952), File("foo1"), "")
	FailIfError(t, err2, "Failed on second create for delete file test")

	_, err = i.CreateRecord(Year(1952), File("foo2"), "")
	FailIfError(t, err, "Failed on third create for delete file test")

	Assert(t, i.ContainsFile(File("foo1")), "File exists before delete")
	FailIfError(t, i.DeleteRecord(r), "Delete should work")
	Assert(t, !i.ContainsFile(File("foo1")), "File removed")
}

func TestCantDeleteNonExistentFile(t *testing.T) {
	i := EmptyIndex()
	r, err := i.CreateRecord(Year(1952), File("foo"), "")
	FailIfError(t, err, "Fail on first create for cant delete non-existent file")

	Assert(t, i.ContainsFile(r.File()), "File does exist before delete.")
	FailIfError(t, i.DeleteRecord(r), "Failed on first delete for delete non existent file")
	Assert(t, !i.ContainsFile(r.File()), "File does not exist before delete.")
	FailIfNoError(t, i.DeleteRecord(r), "Should not be able to delete existant file")
}

func UpdateTags(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	r, err := i.CreateRecord(Year(1952), f, "")
	FailIfError(t, err, "Fail on first update tags create")

	AssertTags(t, r.Tags, []Tag{}, "Empty tags to start")

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
	r, err := i.CreateRecord(Year(1952), f, "note 1")
	FailIfError(t, err, "Failed on first create for update note")

	Assert(t, r.Note == Note("note 1"), "note is correct on construction")
	r.Note = Note("note 2")
	Assert(t, r.Note == Note("note 2"), "note is updated")
}

func TestSwapYear(t *testing.T) {
	i := EmptyIndex()
	f := File("foo")
	r, err := i.CreateRecord(Year(1952), f, "")
	FailIfError(t, err, "Failed on first create for swap year")

	Assert(t, r.Year == Year(1952), "year is correct on construction")
	r.Year = Year(1953)
	Assert(t, r.Year == Year(1953), "year is updated")
}
