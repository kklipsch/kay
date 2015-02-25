package wd

import "os"

type WorkingDirectory string

func Get() (WorkingDirectory, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return WorkingDirectory(""), err
	}

	return WorkingDirectory(pwd), nil
}
