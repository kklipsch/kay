package index

import (
	"time"

	"github.com/kklipsch/kay/chapter"
)

type Year int
type Tag string
type Note string

type Record struct {
	Year        Year
	Note        Note
	DateAdded   time.Time
	LastWritten time.Time
	Tags        []Tag
}

const EmptyYear = Year(0)

func NewRecord(year Year, note Note, tags ...Tag) *Record {
	return &Record{year, note, time.Now(), time.Now(), tags}
}

type Index interface {
	AddChapter(chap chapter.Chapter, record *Record) (*Record, error)
	ContainsChapter(chap chapter.Chapter) bool
	GetRecord(chap chapter.Chapter) (*Record, error)
}
