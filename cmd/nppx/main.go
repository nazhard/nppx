package main

import (
	"fmt"
	"os"

	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	s := setup.InitFunc()

	if s == true {
		exec()
	}
}

func exec() {
	app := &cli.App{
		Name:     "nppx",
		Version:  "experimental-test",
		Usage:    "an tool to manage javascript packages",
		Commands: cmd.Commands(),
		Suggest:  true,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
