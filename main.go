package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
			Action: inKayDir(add),
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
			Action: inKayDir(info),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "mode, m",
					Usage: "info supports several modes (normal, json, year, tags, note, added, written). ",
				},
			},
		},
		{
			Name:   "stat",
			Usage:  "See stats on the current kay directory.",
			Action: inKayDir(stat),
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

	return commands.Initialize(pwd)
}

func add(context *cli.Context, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	y := index.EmptyYear
	if context.IsSet("year") {
		y = index.Year(context.Int("year"))
	}

	t := []index.Tag{}
	if context.IsSet("tags") {
		for _, tag := range strings.Split(context.String("tags"), ",") {
			t = append(t, index.Tag(tag))
		}
	}

	return commands.Add(commands.AddArguments{toChapters(context), y, t}, kd, working)
}

func stat(context *cli.Context, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	return commands.Stat(kd, working)
}

func info(context *cli.Context, kd kaydir.KayDir, working wd.WorkingDirectory) error {
	mode := ""
	if context.IsSet("mode") {
		mode = context.String("mode")
	}

	chapters := toChapters(context)
	if len(chapters) != 1 {
		return fmt.Errorf("info works on 1 and only 1 chapter.")
	}

	return commands.Info(chapters[0], mode, kd, working)
}

func toChapters(context *cli.Context) []chapter.Chapter {
	var chapters []chapter.Chapter
	for _, arg := range context.Args() {
		chapters = append(chapters, chapter.Chapter(arg))
	}
	return chapters
}

func inKayDir(cmd func(*cli.Context, kaydir.KayDir, wd.WorkingDirectory) error) func(*cli.Context) error {

	return func(context *cli.Context) error {

		working, err := wd.Get()
		if err != nil {
			return err
		}

		kd, kdErr := kaydir.Get(working)
		if kdErr != nil {
			return kdErr
		}

		return cmd(context, kd, working)
	}
}
