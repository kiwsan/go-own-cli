package main

import (
	"fmt"
	"log"
	"os"

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
			Name:        "product",
			Aliases:     []string{"p"},
			Usage:       "Shows a list of products or help for a product",
			Description: "This product provider",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"ls"},
					Usage:       "display products list",
					Description: " ",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "Name of the product to show",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("name") == "" {
							fmt.Println("Out of store.")
						} else {
							fmt.Println("This is,", c.String("name"))
						}

						return nil
					},
				},
			},
		},
		{
			Name:        "cart",
			Aliases:     []string{"c"},
			Usage:       "Shows a cart",
			Description: "This cart provider",
			Action: func(c *cli.Context) error {

				fmt.Println("My cart.")

				return nil
			},
		},
		{
			Name:        "order",
			Aliases:     []string{"r"},
			Usage:       "Shows a list of orders or help for an order",
			Description: "This order provider",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"ls"},
					Usage:       "display orders list",
					Description: " ",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Value: "",
							Usage: "id of the order to show",
						},
					},
					Action: func(c *cli.Context) error {
						if c.String("id") == "" {
							fmt.Println("Empty!")
						} else {
							fmt.Println("This is,", c.String("id"))
						}

						return nil
					},
				},
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

