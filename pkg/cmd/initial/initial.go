package initial

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nazhard/nppx"
	"github.com/urfave/cli/v2"
)

type Pkg struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description,omitempty"`
	Main            string            `json:"main,omitempty"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Keywords        []string          `json:"keywords,omitempty"`
	Author          string            `json:"author,omitempty"`
	License         string            `json:"license,omitempty"`
}

func Action(c *cli.Context) error {
	err := nppx.NoPackageJson()
	if err != nil {
		err := nppx.CurrentDirNotEmpty()
		if err != nil {
			fmt.Println(err)
		} else {
			InitPackageJson()
		}
	}

	return nil
}

func InitPackageJson() {
	pwd, _ := os.Getwd()
	pkg := Pkg{
		Name:        "",
		Version:     "0.0.1",
		Description: "",
		Main:        "index.js",
		Scripts: map[string]string{
			"test": "echo \"test\"",
		},
		Dependencies:    map[string]string{},
		DevDependencies: map[string]string{},
		Author:          "",
		License:         "ISC",
	}

	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	file, err := os.Create("package.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	f, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Wrote to " + pwd + "/package.json")
	fmt.Println(string(f))
}
