package main

import (
	"fmt"
	"os"

	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/cli/cmd"
	"github.com/urfave/cli/v2"
)

var visibleCommandTemplate = `{{ $cv := offsetCommands .VisibleCommands 5}}{{range .VisibleCommands}}
   {{$s := join .Names ", "}}{{$s}}{{ $sp := subtract $cv (offset $s 3) }}{{ indent $sp ""}}{{wrap .UsageText $cv}}{{end}}`

func main() {
	s := setup.InitFunc()

	if s == true {
		exec()
	}
}

func exec() {
	app := &cli.App{
		Name:            "nppx",
		Version:         "experimental-test",
		Usage:           "an tool to manage javascript packages",
		Commands:        cmd.Commands(),
		HideHelpCommand: true,
		Suggest:         true,
		CustomAppHelpTemplate: `{{.Name}}{{if .Version}} {{.Version}}{{end}}

Usage: {{.HelpName}}{{if .Commands}} command{{end}}{{if .VisibleFlags}} [flags]{{end}}{{if .ArgsUsage}} {{.ArgsUsage}}{{else}}{{end}}
{{if .Commands}}
Commands: {{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.UsageText}}{{end}}{{else}}{{template "visibleCommandTemplate" .}}{{end}}{{end}}

Flags:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
