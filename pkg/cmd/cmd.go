package cmd

import (
	"fmt"

	"github.com/nazhard/nppx/pkg/cmd/install"
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install global packages",
			Action: func(c *cli.Context) error {
				install.Install(c.Args().Get(0))

				return nil
			},
		},
		{
			Name:      "list",
			Aliases:   []string{"ls"},
			Usage:     "list installed packages",
			UsageText: "nppx list, nppx ls",
			Action: func(c *cli.Context) error {
				fmt.Println("You have no packages")
				// list.Print()

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
