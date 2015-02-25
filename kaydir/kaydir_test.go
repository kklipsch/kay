package kaydir

import (
	"os"
	"testing"

	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestGetFailsOnNonExistentKayDir(t *testing.T) {
	tempdir.In("get-fails", func(dir string) {
		if _, err := Get(wd.WorkingDirectory(dir)); err == nil {
			t.Error("Get did not fail on nonexistent kaydir")
		}
	})
}

func TestMakeCreatesDir(t *testing.T) {
	tempdir.In("make-dir", func(dir string) {
		kayDir, _ := Make(wd.WorkingDirectory(dir))
		stat, err := os.Stat(string(kayDir))
		if err != nil {
			t.Errorf("Make did not make: %v", err)
		}

		if !stat.IsDir() {
			t.Errorf("Make did not make a dir: %v", stat)
		}
	})
}

func TestGetWorksIfExists(t *testing.T) {
	tempdir.In("get-works", func(dir string) {
		working := wd.WorkingDirectory(dir)
		made, _ := Make(working)
		gotten, err := Get(working)
		if err != nil {
			t.Errorf("Get failed: %v", err)
		}

		if made != gotten {
			t.Errorf("Not equivalent: %s:%s", made, gotten)
		}
	})
}
