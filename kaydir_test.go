package main

import (
	"testing"

	"github.com/kklipsch/kay/tempdir"
)

func TestMakeKayDir(t *testing.T) {
	tempdir.In("make kay dir", func(dir string) {
		kayDir := KayDir(dir)

		Assert(t, !kayDir.In(), "Not a kay dir to start")
		kayDir.Make()
		Assert(t, kayDir.In(), "In a kay dir after making")
	})
}
