package main

import (
	"fmt"
	"time"
)

type Year int
type File string
type Tag string
type Note string

type Index struct {
	records map[File]*record
}

func EmptyIndex() *Index {
	return &Index{(map[File]*record{})}
}

func (i *Index) ContainsFile(file File) bool {
	_, contains := i.records[file]
	return contains
}

func (i *Index) Get(file File) *record {
	record, _ := i.records[file]
	return record
}

func (i *Index) CreateRecord(year Year, file File, note Note, tags ...Tag) (*record, error) {
	if i.ContainsFile(file) {
		return nil, fmt.Errorf("File %v exists.", file)
	}

	now := time.Now()
	r := record{}
	r.Year = year
	r.file = file
	r.DateAdded = now
	r.Tags = tags
	r.Note = note
	i.records[file] = &r
	return &r, nil
}

func (i *Index) addRecord(r *record, newFile File, replace bool) (*record, error) {
	if i.ContainsFile(newFile) {
		return nil, fmt.Errorf("Attempting to add file to one that already exists in index: %s", newFile)
	}

	if replace {
		delerr := i.DeleteRecord(r)
		if delerr != nil {
			return nil, delerr
		}
	}

	r.file = newFile
	i.records[newFile] = r

	return r, nil
}

func (i *Index) SetRecord(r *record, newFile File) (*record, error) {
	return i.addRecord(r, newFile, false)
}

func (i *Index) MoveRecord(r *record, newFile File) (*record, error) {
	return i.addRecord(r, newFile, true)
}

func (i *Index) DeleteRecord(r *record) error {
	if !i.ContainsFile(r.File()) {
		return fmt.Errorf("File %v is not in index.", r.File())
	}

	delete(i.records, r.File())
	return nil
}

func (i *Index) Records() []*record {
	records := []*record{}
	for _, record := range i.records {
		records = append(records, record)
	}
	return records
}

type record struct {
	//changing the file should be done through the index
	file File

	Year        Year
	Note        Note
	DateAdded   time.Time
	LastWritten time.Time
	Tags        []Tag
}

func (r record) File() File {
	return r.file
}

func (r record) HasTag(tag Tag) bool {
	//golang doesnt have a contains function for arrays?
	for _, t := range r.Tags {
		if t == tag {
			return true
		}
	}

	return false
}
