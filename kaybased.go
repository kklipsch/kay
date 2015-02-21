package main

import (
	"fmt"
	"os"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/index"
)

func KayBased(action func(c *cli.Context, kayDir KayDir, i index.Index) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		kayDir, err := GetKayDir()
		if err != nil {
			return err
		}

		if !kayDir.In() {
			return fmt.Errorf("This is not a kay directory.")
		}

		i, indexErr := index.IndexDirectory(kayDir.Index())
		if indexErr != nil {
			return indexErr
		}

		return action(c, kayDir, i)
	}
}

func GetKayDir() (KayDir, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return KayDir(pwd), nil
}
