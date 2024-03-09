package cmd

import (
	"github.com/nazhard/nppx/pkg/cmd/install"
	"github.com/nazhard/nppx/pkg/cmd/list"
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install global packages",
			Action: func(c *cli.Context) error {
				install.Install(c.Args().Slice())

				return nil
			},
		},
		{
			Name:      "list",
			Aliases:   []string{"ls"},
			Usage:     "list installed packages",
			UsageText: "nppx list, nppx ls",
			Action: func(c *cli.Context) error {
				list.Print(c.Args().Slice())

				return nil
			},
		},
		{
			Name:      "check",
			Aliases:   []string{"ch"},
			Usage:     "check installed packages",
			UsageText: "nppx check <package>",
		},
	}

	return cmds
}
