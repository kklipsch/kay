package main

import (
	"os"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/kaydir"
)

func Initialize(c *cli.Context) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = kaydir.Make(pwd)
	return err
}
