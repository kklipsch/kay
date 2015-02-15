package main

import "testing"

func TestMakeKayDir(t *testing.T) {
	InTempDir(t, "make", func(dir string) {
		kayDir := KayDir(dir)

		in, err := kayDir.In()
		FailIfError(t, err, "Before in failed")
		Assert(t, !in, "Not a kay dir to start")

		FailIfError(t, kayDir.Make(), "Failed at making")

		in, err = kayDir.In()
		FailIfError(t, err, "After In failed")
		Assert(t, in, "In a kay dir after making")
	})
}

func TestKayDirContentFiles(t *testing.T) {
	InTempDir(t, "make", func(dir string) {
		kayDir := KayDir(dir)

		FailIfError(t, kayDir.Make(), "Failed at making for content")

	})
}
