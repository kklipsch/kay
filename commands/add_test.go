package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestParseTags(t *testing.T) {
	test := func(chap chapter.Chapter, y string) {
		if gy, _ := parseTag(chap); gy != index.Tag(y) {
			t.Errorf("Expected %v Got %v: %v", y, gy, chap)
		}
	}

	test(chapter.Chapter("1942.Foo.doc"), "Foo")
	test(chapter.Chapter("1942.doc"), "")
	test(chapter.Chapter("1944.Foo Bar.doc"), "Foo Bar")
	test(chapter.Chapter("1943.Foo Bar.txt"), "Foo Bar")
	test(chapter.Chapter("1954.Boo_Goo Hoo.doc"), "Boo_Goo Hoo")
	test(chapter.Chapter("1956.Moo-Boo.docx"), "Moo-Boo")
	test(chapter.Chapter("1965.Goo.Soo.Boo.docx"), "Goo.Soo.Boo")
}

func TestUnableToParseTags(t *testing.T) {
	test := func(chap chapter.Chapter) {
		if gy, err := parseTag(chap); err != nil || gy != "" {
			t.Errorf("parse errors %v: %v", chap, gy)
		}
	}

	test(chapter.Chapter("1943x.Foo.doc"))
	test(chapter.Chapter("x1943.Foo.doc"))
	test(chapter.Chapter("Foo.1954"))
	test(chapter.Chapter("Foo.1954.doc"))
	test(chapter.Chapter("Foo.doc"))
}

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
	tempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {

		c1 := chapter.Chapter("foo")
		c2 := chapter.Chapter("moo")
		c3 := chapter.Chapter("boo")

		y := index.Year(1942)

		Add(AddArguments{[]chapter.Chapter{c1, c2, c3}, y, []index.Tag{}, index.Note("")}, kd, w)

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

func TestAddExplicitNote(t *testing.T) {
	tempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {

		c1 := chapter.Chapter("foo")
		c2 := chapter.Chapter("moo")
		c3 := chapter.Chapter("boo")

		n := index.Note("noooo")

		Add(AddArguments{[]chapter.Chapter{c1, c2, c3}, index.Year(1942), []index.Tag{}, n}, kd, w)

		i, _ := index.Get(kd)

		rec1, _ := i.GetRecord(c1)
		if rec1.Note != n {
			t.Errorf("rec 1 was incorrect: %v", rec1)
		}

		rec2, _ := i.GetRecord(c2)
		if rec2.Note != n {
			t.Errorf("rec 2 was incorrect: %v", rec2)
		}

		rec3, _ := i.GetRecord(c3)
		if rec3.Note != n {
			t.Errorf("rec 3 was incorrect: %v", rec3)
		}
	})
}

func tagsEqual(rec *index.Record, tags ...string) error {
	t := []string{}
	for _, tag := range rec.Tags {
		t = append(t, string(tag))
	}
	received := strings.Join(t, ",")

	expected := strings.Join(tags, ",")

	if received != expected {
		return fmt.Errorf("Expected %s Received %s", expected, received)
	}
	return nil
}

func TestAddExplicitTags(t *testing.T) {
	tempInit(func(w wd.WorkingDirectory, kd kaydir.KayDir) {

		c1 := chapter.Chapter("foo")
		c2 := chapter.Chapter("moo")
		c3 := chapter.Chapter("boo")

		y := index.Year(1942)

		Add(AddArguments{[]chapter.Chapter{c1, c2, c3}, y, []index.Tag{index.Tag("a"), index.Tag("b")}, index.Note("")}, kd, w)

		i, _ := index.Get(kd)

		rec1, _ := i.GetRecord(c1)
		if err := tagsEqual(rec1, "a", "b"); err != nil {
			t.Errorf("Rec 1 was incorrect: %v", err)
		}

		rec2, _ := i.GetRecord(c2)
		if err := tagsEqual(rec2, "a", "b"); err != nil {
			t.Errorf("rec2  was incorrect: %v", err)
		}

		rec3, _ := i.GetRecord(c3)
		if err := tagsEqual(rec3, "a", "b"); err != nil {
			t.Errorf("rec3  was incorrect: %v", err)
		}
	})
}
