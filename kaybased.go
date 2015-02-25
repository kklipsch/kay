package main

import (
	"os"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
)

func KayBased(action func(c *cli.Context, kd kaydir.KayDir, i index.Index) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		kd, err := GetKayDir()
		if err != nil {
			return err
		}

		i, indexErr := index.IndexDirectory(kd)
		if indexErr != nil {
			return indexErr
		}

		return action(c, kd, i)
	}
}

func GetKayDir() (kaydir.KayDir, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return kaydir.Get(pwd)
}
