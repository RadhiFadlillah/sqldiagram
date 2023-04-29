package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "sqldiagram",
		Usage:     "generate ERD from SQL file(s)",
		UsageText: "sqldiagram [global flags] command [command flags] <input-dir>",
		Commands: []*cli.Command{
			cmdMySql(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
