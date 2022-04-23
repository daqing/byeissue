package main

import (
	"log"
	"os"

	"github.com/daqing/byeissue/db/migration"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "migrate",
		Usage: "run database migrations",
		Action: func(c *cli.Context) error {
			migration.CreateUsersTable()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
