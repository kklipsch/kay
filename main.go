package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kklipsch/cli"
)

func main() {
	err := NewKay().Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func NewKay() *cli.App {
	app := cli.NewApp()
	app.Name = "kay"
	app.Usage = "Highly specific content management system for grandparent biographies"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "years",
			Usage:  "List available years",
			Action: inKay("years"),
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "missing, m",
					Usage: "show missing years instead of existing ones.",
				},
				cli.IntFlag{
					Name:  "from, f",
					Usage: "which year to start on (defaults to first year in index if not provided)",
					Value: 1900,
				},
				cli.IntFlag{
					Name:  "to, t",
					Usage: "which year to end on (defaults to this year if not provided)",
					Value: time.Now().Year(),
				},
			},
		},
		{
			Name:   "tags",
			Usage:  "List available tags",
			Action: inKay("tags"),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "add",
			Usage:  "[files] - add files to an index.",
			Action: KayBased(Add),
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "year, y",
					Usage: "which year the files should be attached to. Will fail if not provided or parsed.",
				},
				cli.StringFlag{
					Name:  "tags, t",
					Usage: "a csv list of tags for the file. Will NOT fail if not provided or parsed.",
				},
			},
		},
		{
			Name:   "rm",
			Usage:  "[files] - rm files from the index.",
			Action: inKay("rm"),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "html",
			Usage:  "create the website",
			Action: inKay("html"),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "txt",
			Usage:  "[files] - show the text (if available) for a given file",
			Action: inKay("txt"),
			Flags:  []cli.Flag{},
		},
		{
			Name:   "stat",
			Usage:  "See stats on the current kay directory.",
			Action: KayBased(Stat),
		},
		{
			Name:   "init",
			Usage:  "initialize a new kay directory",
			Action: Initialize,
		},
	}

	return app
}

func inKay(name string) func(c *cli.Context) error {
	return KayBased(func(c *cli.Context, kayDir KayDir, index *index) error {
		fmt.Printf(name)
		return nil
	})
}
