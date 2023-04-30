package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appFlNoGroup      = "no-group"
	appFlNoGroupAlias = []string{"ng"}
)

func main() {
	app := &cli.App{
		Name:      "sqldiagram",
		Usage:     "generate ERD from SQL file(s)",
		UsageText: "sqldiagram [global options] command [command options] <input-dir>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Value:   false,
				Name:    appFlNoGroup,
				Aliases: appFlNoGroupAlias,
				Usage:   "don't render separate file as group",
			},
		},
		Commands: []*cli.Command{
			cmdMySql(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
