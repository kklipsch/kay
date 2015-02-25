package commands

import (
	"testing"

	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestRunFailsOnNoKayDir(t *testing.T) {
	tempdir.In("run-command-fail", func(dir string) {

		err := RunWithKayDir(nil, wd.WorkingDirectory(dir), func(args Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
			t.Error("Should not have been called")
			return nil
		})

		if err == nil {
			t.Error("No error")
		}
	})
}

func TestRunCommandsWithKayDir(t *testing.T) {
	tempdir.In("run-command-fail", func(dir string) {
		called := false
		working := wd.WorkingDirectory(dir)
		Initialize(nil, working)

		RunWithKayDir(nil, working, func(args Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
			called = true
			return nil
		})

		if !called {
			t.Error("command not run")
		}
	})
}
