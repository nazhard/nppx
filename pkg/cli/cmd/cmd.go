package cmd

import (
	"github.com/nazhard/nppx/pkg/cli/cmd/initial"
	"github.com/nazhard/nppx/pkg/cli/cmd/install"
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:      "install",
			Aliases:   []string{"i"},
			Usage:     "Install all dependencies when no args given.",
			UsageText: "nppx i, nppx i <pkg>",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "dev",
					Aliases: []string{"D"},
					Usage:   "Install devDependencies",
				},
				&cli.BoolFlag{
					Name:    "prod",
					Aliases: []string{"P"},
					Usage:   "Packages in devDependencies won't be installed",
				},
			},
			Category: "\n  Manage your dependencies",
			Action:   install.Action,
		},
		{
			Name:    "update",
			Aliases: []string{"up"},
			Usage:   "Update all dependencies",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "g",
					Usage: "Update everything",
				},
			},
			Category: "\n  Manage your dependencies",
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "Removes packages from the project's package.json",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "g",
					Usage: "Remove dependencies from entire system",
				},
			},
			Category: "\n  Manage your dependencies",
		},
		{
			Name:     "init",
			Usage:    "Initialize simple package.json",
			Action:   initial.Action,
			Category: "\n  Others",
		},
	}

	return cmds
}
