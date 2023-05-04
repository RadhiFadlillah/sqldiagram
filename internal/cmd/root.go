package cmd

import (
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:      "sqldiagram",
		Usage:     "generate ERD from SQL file(s) as SVG file",
		UsageText: "sqldiagram command [command options] <input-1> ... <input-n>",
		Commands:  []*cli.Command{cmdMySql},
	}
}
