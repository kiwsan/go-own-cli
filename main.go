package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Simple kiwsa CLI"
	app.Usage = "An example CLI for learning"
	app.Author = "kiwsa"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
