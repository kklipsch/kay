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

		c := chapter.Chapter("foo")

		Add(Arguments{[]chapter.Chapter{c}, index.Year(1941)}, kd, w)

		i, _ := index.Get(kd)
		if !i.ContainsChapter(c) {
			t.Error("Add with a year argument did not add year.")
		}
	})
}
