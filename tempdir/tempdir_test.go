package tempdir

import (
	"os"
	"strings"
	"testing"

	"github.com/kklipsch/kay/wd"
)

func TestTempDirCreatesDirectory(t *testing.T) {
	In(func(dir string) {
		stat, err := os.Stat(dir)
		if err != nil {
			t.Errorf("Error on stats of temp dir: %v", err)
		}

		if !stat.IsDir() {
			t.Errorf("Tempdir is not a directroy: %v", stat)
		}
	})
}

func TestTempDirDeletesItself(t *testing.T) {
	var created string
	In(func(dir string) {
		created = dir
	})

	_, err := os.Stat(created)
	if err == nil {
		t.Errorf("File exists.")
	}
}

func TestTempWd(t *testing.T) {
	called := false
	err := TempWd(func(working wd.WorkingDirectory) {
		called = true
	})

	if err != nil {
		t.Error("Failed: %v", err)
	}

	if !called {
		t.Errorf("Wasn't called")
	}
}

func TestTempWdHasNameInIt(t *testing.T) {
	TempWd(func(working wd.WorkingDirectory) {
		if !strings.Contains(string(working), "tempdir_test.go") {
			t.Errorf("Does not have call location: %v", working)
		}
	})
}
