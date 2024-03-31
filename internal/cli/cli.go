package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Command represents a CLI command
type Command struct {
	Name        string
	Usage       string
	Help        string
	Description string
	Aliases     []string
	Action      func()
}

// CLI represents the command line interface
type CLI struct {
	Name     string
	Usage    string
	Version  string
	commands map[string]Command
}

// New creates a new CLI instance
func New() *CLI {
	return &CLI{
		Name:     "",
		Version:  "",
		Usage:    "",
		commands: make(map[string]Command),
	}
}

// AddCommand registers a new command
func (c *CLI) AddCommand(cmd Command) {
	c.commands[cmd.Name] = cmd
	for _, alias := range cmd.Aliases {
		c.commands[alias] = cmd
	}
}

// ExecuteCommand executes a command by name
func (c *CLI) ExecuteCommand(name string, args []string) {
	cmd, ok := c.commands[name]
	if !ok {
		fmt.Println("Unknown command:", name)
		c.PrintUsage()
		return
	}

	// Check if help flag is provided
	if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
		c.PrintCommandHelp(name)
		return
	}

	cmd.Action()
}

// PrintUsage prints usage information for the CLI program
func (c *CLI) PrintUsage() {
	fmt.Printf("Usage: %s %s \n\n", c.Name, c.Usage)
	fmt.Println("Commands:")
	printedCommands := make(map[string]bool)
	for _, cmd := range c.commands {
		if !printedCommands[cmd.Name] {
			fmt.Printf("  %s\t%s\n", cmd.Name, cmd.Usage)
			printedCommands[cmd.Name] = true
		}

		// I don't know but this is to circumvent commands that are printed twice
		for _, alias := range cmd.Aliases {
			if !printedCommands[alias] {
				fmt.Print()
				printedCommands[alias] = true
			}
		}
	}
}

// PrintCommandHelp  prints help information for a specific command
func (c *CLI) PrintCommandHelp(name string) {
	cmd, ok := c.commands[name]
	if !ok {
		fmt.Println("Unknown command:", name)
		c.PrintUsage()
		return
	}

	tmpl := `Usage: {{.CLIName}} {{.CmdName}} {{.Usage}}
{{if .Aliases}}Aliases: {{join .Aliases ", "}}{{end}}
{{if .Description}}{{.Description}}{{end}}
{{.Help}}`

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	helpTemplate, err := template.New("help").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing help template:", err)
		return
	}

	err = helpTemplate.Execute(os.Stdout, map[string]interface{}{
		"CLIName":     c.Name,
		"CmdName":     cmd.Name,
		"Usage":       cmd.Usage,
		"Aliases":     cmd.Aliases,
		"Description": cmd.Description,
		"Help":        cmd.Help,
	})
	if err != nil {
		fmt.Println("Error executing help template:", err)
		return
	}
}

//=====
// THIS IS USAGE FOR TEST PURPOSE. FOR NOW I HAVE NO IDEA TO WRITE TESTING UNIT BECAUSE IT IS STILL UNDER DEVELOPMENT
//=====
// func main() {
// 	cli := New()
//	cli.Name = "neko"

// 	echoCmd := Command{
// 		Name:    "echo",
// 		Usage:   "[-m <message>]",
// 		Aliases: []string{"say"},
// 		Help:    "Echo!",
// 		Action: func() {
// 			fmt.Println("Echo cho ho")
// 		},
// 	}
// 	nx := Command{
// 		Name:    "waifu",
// 		Aliases: []string{"xxx"},
// 		Help:    "Download trojan",
// 		Action: func() {
//			fmt.Println("Haik, onii-chan")
// 		},
// 	}
// 	cli.AddCommand(echoCmd)
// 	cli.AddCommand(nx)

// 	// Parse command line arguments
// 	if len(os.Args) < 2 {
// 		cli.PrintUsage()
// 		return
// 	}

// 	// Execute the specified command
// 	cmdName := os.Args[1]
// 	cli.ExecuteCommand(cmdName, os.Args[2:])
// }
