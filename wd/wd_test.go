package wd

import (
	"os"
	"testing"
)

func TestGetWd(t *testing.T) {
	working, _ := Get()
	pwd, _ := os.Getwd()

	if pwd != string(working) {
		t.Errorf("%s != %s", pwd, working)
	}
}
