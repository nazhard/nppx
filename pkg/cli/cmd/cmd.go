package cmd

import (
	"github.com/nazhard/nppx/pkg/cli/cmd/initial"
	"github.com/nazhard/nppx/pkg/cli/cmd/install"
	"github.com/urfave/cli/v2"
)

var help = `Usage: {{.HelpName}} {{if .Usage}}{{.Usage}}{{end}}[options]
{{if .Aliases}}
Alias: {{range .Aliases}}{{.}}{{end}}
{{end}}
{{.Description}}

Options:
	{{range .Flags}}{{.}}
	{{end}}`

func Commands() []*cli.Command {
	cmds := []*cli.Command{
		{
			Name:        "install",
			Aliases:     []string{"i"},
			Description: "Installs all dependencies of the project in the current working directory.",
			UsageText:   "Install all dependencies for a project",
			Category:    "\n  Manage your dependencies",
			Action:      install.Action,
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
			CustomHelpTemplate: help,
		},
		{
			Name:        "update",
			Aliases:     []string{"up"},
			Description: "Upadate all dependencies",
			UsageText:   "Update all dependencies",
			Category:    "\n  Manage your dependencies",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "g",
					Usage: "Update everything",
				},
			},
			CustomHelpTemplate: help,
		},
		{
			Name:        "remove",
			Aliases:     []string{"rm"},
			Description: "Removes packages from the project's package.json",
			UsageText:   "Removes packages from the project's package.json",
			Category:    "\n  Manage your dependencies",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "g",
					Usage: "Remove dependencies from entire system",
				},
			},
			CustomHelpTemplate: help,
		},
		{
			Name:               "init",
			Description:        "Initialize simple package.json",
			UsageText:          "Initialize simple package.json",
			Action:             initial.Action,
			Category:           "\n  Others",
			CustomHelpTemplate: help,
		},
	}

	return cmds
}
