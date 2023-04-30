package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appFlNoGroup      = "no-group"
	appFlNoGroupAlias = []string{"ng"}

	appFlRawD2      = "raw-d2"
	appFlRawD2Alias = []string{"raw"}
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
			&cli.BoolFlag{
				Value:   false,
				Name:    appFlRawD2,
				Aliases: appFlRawD2Alias,
				Usage:   "render as raw D2 scripts",
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
