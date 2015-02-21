package index

import "time"

type Year int
type File string
type Tag string
type Note string

type Record struct {
	Year        Year
	Note        Note
	DateAdded   time.Time
	LastWritten time.Time
	Tags        []Tag
}

func NewRecord(year Year, note Note, tags ...Tag) *Record {
	return &Record{year, note, time.Now(), time.Now(), tags}
}

type Index interface {
	AddRecord(file File, record *Record) (*Record, error)
	ContainsFile(file File) bool
}
