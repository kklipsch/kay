package kayignore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/wd"
)

func TestExplicitMatch(t *testing.T) {
	ki := getPatterns("foo")

	if ki.Ignore(chapter.Chapter("bar")) {
		t.Errorf("Should not ignored")
	}

	if !ki.Ignore(chapter.Chapter("foo")) {
		t.Errorf("Should have ignored")
	}
}

func TestPatternMatch(t *testing.T) {
	ki := getPatterns("^bar*")

	if ki.Ignore(chapter.Chapter("foo")) {
		t.Errorf("Should not have ignored")
	}

	if ki.Ignore(chapter.Chapter("good barbque")) {
		t.Errorf("Should not have ignored")
	}

	if !ki.Ignore(chapter.Chapter("barbque")) {
		t.Errorf("Should have ignored")
	}
}

func TestNoFileGetIgnore(t *testing.T) {
	tmpdir, _ := ioutil.TempDir("", "kayignore-noload")
	defer os.RemoveAll(tmpdir)

	ki, _ := Get(wd.WorkingDirectory(tmpdir))

	if ki.Ignore("foo") {
		t.Errorf("Shouldnt fail")
	}
}

func TestFileLoadGetIgnore(t *testing.T) {
	tmpdir, _ := ioutil.TempDir("", "kayignore-noload")
	defer os.RemoveAll(tmpdir)

	ignore := fmt.Sprintf("foo\n.bar.\n\n\n    \ngoo")
	ioutil.WriteFile(filepath.Join(tmpdir, ignoreFile), []byte(ignore), 0744)

	ki, _ := Get(wd.WorkingDirectory(tmpdir))
	if ki.Ignore("boogey") {
		t.Errorf("Should not ignore")
	}

	if !ki.Ignore(ignoreFile) {
		t.Errorf("Should ignore the ignore file implicitly")
	}

	if !ki.Ignore("foo") {
		t.Errorf("Should ignore")
	}

	if !ki.Ignore("good barbq") {
		t.Errorf("Should ignore")
	}

	if !ki.Ignore("goo") {
		t.Errorf("Should ignore")
	}
}
