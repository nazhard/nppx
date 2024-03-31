package initial

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nazhard/nppx/internal/resolver"
	"github.com/nazhard/nppx/internal/utils/check"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	err := check.NoPackageJson()
	if err != nil {
		err := check.CurrentDirNotEmpty()
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
	pkg := resolver.Pkg{
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

	fmt.Println("Wrote to " + pwd + "/package.json \n")
	fmt.Println(string(f))
}
