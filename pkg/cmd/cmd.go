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
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "D",
					Usage: "write to devDependencies",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Bool("D") == true {
					install.Install(c.Args().Slice(), true)
				} else {
					install.Install(c.Args().Slice(), false)
				}
				return nil
			},
		},
		{
			Name:      "list",
			Aliases:   []string{"ls"},
			Usage:     "list installed packages",
			UsageText: "nppx list, nppx ls",
			Action: func(c *cli.Context) error {
				list.Print()

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
