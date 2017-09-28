package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/whatdacode/GoREST/database"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "options for database",
			Subcommands: []cli.Command{
				{
					Name:  "migrate",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						database.Migrations()
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
