package wd

import (
	"os"
	"path/filepath"
)

//WorkingDirectory is the directory where kay is invoked
type WorkingDirectory string

//Get returns the WorkingDirectory
func Get() (WorkingDirectory, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return WorkingDirectory(""), err
	}

	return WorkingDirectory(pwd), nil
}

//Path combines a given path with the WorkingDirectory
func (working WorkingDirectory) Path(rest string) string {
	return filepath.Join(string(working), rest)
}
