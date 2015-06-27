package index

import (
	"time"

	"github.com/kklipsch/kay/chapter"
)

//Year is the year of the chapter.
type Year int

//Tag is a string meta information, multiples are allowed per chapter.
type Tag string

//Note is a user entered note, only 1 is allowed per chapter.
type Note string

//Record represents all of the metadata stored in the index for a chapter.
type Record struct {
	Year        Year
	Note        Note
	DateAdded   time.Time
	LastWritten time.Time
	Tags        []Tag
}

//EmptyYear is 0
const EmptyYear = Year(0)

//NewRecord creates a record with the passed in values and UTC Now for the times.
func NewRecord(year Year, note Note, tags ...Tag) *Record {
	return &Record{year, note, time.Now().UTC(), time.Now().UTC(), tags}
}

//Index is an abstraction around manipulating the records in the .kay directory in memory
type Index interface {
	AddChapter(chap chapter.Chapter, record *Record) (*Record, error)
	ContainsChapter(chap chapter.Chapter) bool
	GetRecord(chap chapter.Chapter) (*Record, error)
	AllIndexed() []chapter.Chapter
}
