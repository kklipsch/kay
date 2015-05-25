package commands

import (
	"testing"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestParseYear(t *testing.T) {
	test := func(chap chapter.Chapter, y int) {
		if gy, _ := parseYear(chap); gy != index.Year(y) {
			t.Errorf("Expected %v Got %v: %v", y, gy, chap)
		}
	}

	test(chapter.Chapter("1942.Foo.doc"), 1942)
	test(chapter.Chapter("1944.Foo Bar.doc"), 1944)
	test(chapter.Chapter("1943.Foo Bar.txt"), 1943)
	test(chapter.Chapter("1954.Boo_Goo Hoo.doc"), 1954)
	test(chapter.Chapter("1956.Moo-Boo.docx"), 1956)
	test(chapter.Chapter("1965.Goo.Soo.Boo.docx"), 1965)
}

func TestUnableToParse(t *testing.T) {
	test := func(chap chapter.Chapter) {
		if gy, err := parseYear(chap); err == nil {
			t.Errorf("Should have errors %v: %v", chap, gy)
		}
	}

	test(chapter.Chapter("1943x.Foo.doc"))
	test(chapter.Chapter("x1943.Foo.doc"))
	test(chapter.Chapter("Foo.1954.doc"))
	test(chapter.Chapter("Foo.doc"))
}

func TestFailureIfNoIndex(t *testing.T) {
	tempdir.TempWd(func(w wd.WorkingDirectory) {
		k := kaydir.KayDir(string(w))
		if err := Add(AddArguments{}, k, w); err == nil {
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

		Add(AddArguments{[]chapter.Chapter{c1, c2, c3}, y}, kd, w)

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
