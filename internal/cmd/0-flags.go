package cmd

import "github.com/urfave/cli/v2"

var (
	fNoGroup   = "no-group"
	fRawD2     = "raw-d2"
	fOutput    = "output"
	fDirection = "direction"

	faNoGroup   = []string{"ng"}
	faRawD2     = []string{"raw"}
	faOutput    = []string{"o"}
	faDirection = []string{"dir"}
)

func initFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Value:   false,
			Name:    fNoGroup,
			Aliases: faNoGroup,
			Usage:   "don't render separate file as group",
		},
		&cli.BoolFlag{
			Value:   false,
			Name:    fRawD2,
			Aliases: faRawD2,
			Usage:   "render as raw D2 scripts",
		},
		&cli.StringFlag{
			Name:    fOutput,
			Aliases: faOutput,
			Usage:   "write to specified path (if empty will use stdout)",
		},
		&cli.StringFlag{
			Name:    fDirection,
			Aliases: faDirection,
			Usage:   "direction of chart (up|down|right|left, default right)",
		},
	}
}
