package main

import (
	"os"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
)

func Initialize(c *cli.Context) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	kd, makeErr := kaydir.Make(pwd)
	if makeErr != nil {
		return err
	}

	return index.Make(kd)
}
