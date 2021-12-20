package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	ss := cli.StringSlice{}

	app := &cli.App{
		Name:  "mdbox",
		Usage: "markdown utilities",
		Commands: []*cli.Command{
			{
				Name:  "mv",
				Usage: "move a file somewhere else, updates backlinks",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:        "exts",
						Destination: &ss,
						Required:    false,
						Value:       cli.NewStringSlice(".markdown", ".md"),
					},
				},
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()
					if len(args) < 2 {
						return fmt.Errorf("wrong number of arguments, expected <source> <destination>")
					}
					pwd, err := os.Getwd()
					if err != nil {
						log.Fatal("error getting current directory")
					}

					return mv(mvParams{
						pwd:      pwd,
						src:      args[0],
						dst:      args[1],
						fileExts: ss.Value(),
					})
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
