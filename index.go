package main

import (
	"fmt"
	"time"
)

type Year int
type File string
type Tag string
type Note string

type index struct {
	records map[File]*record
}

func EmptyIndex() *index {
	return &index{(map[File]*record{})}
}

func (i *index) ContainsFile(file File) bool {
	_, contains := i.records[file]
	return contains
}

func (i *index) Get(file File) *record {
	record, _ := i.records[file]
	return record
}

func (i *index) CreateRecord(year Year, file File, note Note, tags ...Tag) (*record, error) {
	r := record{file, year, note, time.Now(), time.Now(), tags}
	return i.addRecord(&r, file, false)
}

func (i *index) addRecord(r *record, newFile File, replace bool) (*record, error) {
	if i.ContainsFile(newFile) {
		return nil, fmt.Errorf("Attempting to add record for file that exists in index: %s", newFile)
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

func (i *index) SetRecord(r *record, newFile File) (*record, error) {
	return i.addRecord(r, newFile, false)
}

func (i *index) MoveRecord(r *record, newFile File) (*record, error) {
	return i.addRecord(r, newFile, true)
}

func (i *index) DeleteRecord(r *record) error {
	if !i.ContainsFile(r.File()) {
		return fmt.Errorf("File %v is not in index.", r.File())
	}

	delete(i.records, r.File())
	return nil
}

func (i *index) Records() []*record {
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
