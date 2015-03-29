package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kklipsch/cli"
	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/commands"
	"github.com/kklipsch/kay/index"
	"github.com/kklipsch/kay/kaydir"
	"github.com/kklipsch/kay/wd"
)

func main() {
	err := NewKay().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
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
			Action: KayBased(commands.Add),
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
			Action: KayBased(commands.Stat),
		},
		{
			Name:   "init",
			Usage:  "initialize a new kay directory",
			Action: Init,
		},
	}

	return app
}

func Init(context *cli.Context) error {
	pwd, err := wd.Get()
	if err != nil {
		return err
	}

	return commands.Initialize(toArguments(context), pwd)
}

func KayBased(cmd func(commands.Arguments, kaydir.KayDir, wd.WorkingDirectory) error) func(*cli.Context) error {
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
	return KayBased(func(args commands.Arguments, kd kaydir.KayDir, working wd.WorkingDirectory) error {
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
	chapters := make([]chapter.Chapter, 0)
	for _, arg := range context.Args() {
		chapters = append(chapters, chapter.Chapter(arg))
	}
	return chapters
}
