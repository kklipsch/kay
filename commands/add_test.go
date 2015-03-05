package commands

import (
	"testing"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestFailureIfNoIndex(t *testing.T) {
	tempdir.TempWd(func(w wd.WorkingDirectory) {
		k := kaydir.KayDir(string(w))
		if err := Add(Arguments{}, k, w); err == nil {
			t.Errorf("Did not fail")
		}
	})
}

func TestAddExplicitYear(t *testing.T) {
	TempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {

		c1 := chapter.Chapter("foo")
		c2 := chapter.Chapter("moo")
		c3 := chapter.Chapter("boo")

		y := index.Year(1942)

		Add(Arguments{[]chapter.Chapter{c1, c2, c3}, y}, kd, w)

		i, _ := index.Get(kd)

		rec1, _ := i.GetRecord(c1)
		if rec1.Year != y {
			t.Errorf("Year 1 was incorrect: %v", rec1)
		}

		rec2, _ := i.GetRecord(c2)
		if rec2.Year != y {
			t.Errorf("Year 2 was incorrect: %v", rec2)
		}

		rec3, _ := i.GetRecord(c3)
		if rec3.Year != y {
			t.Errorf("Year 3 was incorrect: %v", rec3)
		}
	})
}
