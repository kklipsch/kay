package main

import (
	"fmt"
	"os"

	"github.com/kklipsch/cli"
)

func KayBased(action func(c *cli.Context, kayDir KayDir, index *Index) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		kayDir, err := GetKayDir()
		if err != nil {
			return err
		}

		in, kayDirErr := kayDir.In()
		if kayDirErr != nil {
			return kayDirErr
		}

		if !in {
			return fmt.Errorf("This is not a kay directory.")
		}

		index, indexerr := BuildIndex(kayDir.Index())
		if indexerr != nil {
			return indexerr
		}

		actionerr := action(c, kayDir, index)
		if actionerr != nil {
			return actionerr
		}

		return WriteIndex(kayDir.Index(), index)
	}
}

func GetKayDir() (KayDir, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return KayDir(pwd), nil
}
