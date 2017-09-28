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
					Usage: "migrate the database",
					Action: func(c *cli.Context) error {
						database.Migrate()
						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "rollback the last migration",
					Action: func(c *cli.Context) error {
						database.Rollback()
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
