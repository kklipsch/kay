package wd

import (
	"os"
	"path/filepath"
)

type WorkingDirectory string

func Get() (WorkingDirectory, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return WorkingDirectory(""), err
	}

	return WorkingDirectory(pwd), nil
}

func (working WorkingDirectory) Path(rest string) string {
	return filepath.Join(string(working), rest)
}
