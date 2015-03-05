package commands

import (
	"testing"

	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/tempdir"
	"github.com/kklipsch/kay/wd"
)

func TestRunFailsOnNoKayDir(t *testing.T) {
	tempdir.TempWd(func(dir wd.WorkingDirectory) {

		err := RunWithKayDir(Arguments{}, dir, func(args Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
			t.Error("Should not have been called")
			return nil
		})

		if err == nil {
			t.Error("No error")
		}
	})
}

func TestRunCommandsWithKayDir(t *testing.T) {
	tempdir.TempWd(func(working wd.WorkingDirectory) {
		called := false
		Initialize(Arguments{}, working)

		RunWithKayDir(Arguments{}, working, func(args Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
			called = true
			return nil
		})

		if !called {
			t.Error("command not run")
		}
	})
}
