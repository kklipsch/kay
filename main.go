package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/commands"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func main() {
	err := newKay().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newKay() *cli.App {
	app := cli.NewApp()
	app.Name = "kay"
	app.Usage = "Highly specific content management system for grandparent biographies"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "add",
			Usage:  "[files] - add files to an index.",
			Action: kayBased(commands.Add),
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
			Name:   "info",
			Usage:  "[file] - info displays metadata for a file.",
			Action: kayBased(commands.Info),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "mode, m",
					Usage: "info supports several modes (json, year, tags, notes, added, written). ",
				},
			},
		},
		{
			Name:   "stat",
			Usage:  "See stats on the current kay directory.",
			Action: kayBased(commands.Stat),
		},
		{
			Name:   "init",
			Usage:  "initialize a new kay directory",
			Action: kayInit,
		},
	}

	return app
}

func kayInit(context *cli.Context) error {
	pwd, err := wd.Get()
	if err != nil {
		return err
	}

	return commands.Initialize(toArguments(context), pwd)
}

func kayBased(cmd func(commands.Arguments, kaydir.KayDir, wd.WorkingDirectory) error) func(*cli.Context) error {
	return func(context *cli.Context) error {
		pwd, err := wd.Get()
		if err != nil {
			return err
		}

		args := toArguments(context)
		return commands.RunWithKayDir(args, pwd, cmd)
	}
}

func inKay(name string) func(c *cli.Context) error {
	return kayBased(func(args commands.Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
		fmt.Printf(name)
		return nil
	})
}

func toArguments(c *cli.Context) commands.Arguments {
	y := index.EmptyYear
	if c.IsSet("year") {
		y = index.Year(c.Int("year"))
	}

	return commands.Arguments{toChapters(c), y}
}

func toChapters(context *cli.Context) []chapter.Chapter {
	var chapters []chapter.Chapter
	for _, arg := range context.Args() {
		chapters = append(chapters, chapter.Chapter(arg))
	}
	return chapters
}
