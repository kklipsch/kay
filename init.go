package main

import "github.com/kklipsch/cli"

func Initialize(c *cli.Context) error {
	init, err := GetKayDir()
	if err != nil {
		return err
	}

	return init.Make()
}
