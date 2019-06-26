package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var message = []string{"Hello world."}

func info() {
	app.Name = "Own CLI"
	app.Usage = "An example CLI"
	app.Author = "kiwsan"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "products",
			Aliases: []string{"p"},
			Usage:   "Show products",
			Action: func(c *cli.Context) {
				pe := "products"
				products := append(message, pe)
				m := strings.Join(products, " ")
				fmt.Println(m)
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
