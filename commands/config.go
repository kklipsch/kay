package commands

import (
	"fmt"

	"github.com/kklipsch/kay/config"
	"github.com/kklipsch/kay/kaydir"
)

//GetConfig prints either the config or the specified variable.
func GetConfig(kd kaydir.KayDir, variable string) error {
	config, err := config.Get(kd)
	if err != nil {
		return err
	}

	if variable == "" {
		fmt.Printf("%v\n", config)
		return nil
	}

	switch variable {
	case "first":
		fmt.Printf("%s\n", config.First)
	case "last":
		fmt.Printf("%s\n", config.Last)
	default:
		fmt.Printf("%v\n", config)
	}

	return nil
}
