package main

import "testing"

func TestMakeKayDir(t *testing.T) {
	InTempDir(t, "make", func(dir string) {
		kayDir := KayDir(dir)

		Assert(t, !kayDir.In(), "Not a kay dir to start")
		kayDir.Make()
		Assert(t, kayDir.In(), "In a kay dir after making")
	})
}

func TestKayDirContentFiles(t *testing.T) {
	InTempDir(t, "make", func(dir string) {
		kayDir := KayDir(dir)

		FailIfError(t, kayDir.Make(), "Failed at making for content")

	})
}
