package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	appNoGroup = "no-group"
	appRawD2   = "raw-d2"
	appOutput  = "output"

	appAliasNoGroup = []string{"ng"}
	appAliasRawD2   = []string{"raw"}
	appAliasOutput  = []string{"o"}
)

func main() {
	app := &cli.App{
		Name:      "sqldiagram",
		Usage:     "generate ERD from SQL file(s) as SVG file",
		UsageText: "sqldiagram [global options] command [command options] <input-1> ... <input-n>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Value:   false,
				Name:    appNoGroup,
				Aliases: appAliasNoGroup,
				Usage:   "don't render separate file as group",
			},
			&cli.BoolFlag{
				Value:   false,
				Name:    appRawD2,
				Aliases: appAliasRawD2,
				Usage:   "render as raw D2 scripts",
			},
			&cli.StringFlag{
				Name:    appOutput,
				Aliases: appAliasOutput,
				Usage:   "write to specified path (if empty will use stdout)",
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
