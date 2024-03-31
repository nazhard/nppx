package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/nazhard/nppx/internal/lockfile"
	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/internal/utils/check"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	// Checking some necessary dirs and files
	_, err := os.Stat(".nppx")
	if os.IsNotExist(err) {
		_ = os.Mkdir(".nppx", os.ModePerm)
		setup.CreateDotNPPXDir()
	}
	_, err = os.Stat("nppx-lock.json")
	if os.IsNotExist(err) {
		lockfile.GenerateLockfile()
	}

	if c.Bool("dev") {
		install(c.Args().Slice(), "dev")
	} else if c.Bool("prod") {
		install(c.Args().Slice(), "prod")
	} else {
		install(c.Args().Slice(), "def")
	}
	return nil
}

func install(args []string, action string) {
	err := check.NoPackageJson()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err == nil {
		if len(args) == 0 {
			name, version, exist := checkCache()
			if exist == true {
				installFromCache(name, version, action)
			} else {
				do(args, action, false)
			}
		}

		if len(args) != 0 {
			name, version, x := checkCacheWithArgs(args)
			if x == true {
				installFromCache(name, version, action)
			} else {
				do(args, action, true)
			}
		}

		// err = fs.CreateSymlinks(".nppx/.modules.json")
	}
}

func contains(s, c, cc string) (string, string, bool) {
	if strings.Contains(s, c) && !strings.Contains(s, cc) {
		p := strings.Split(s, c)
		if len(p) == 2 {
			name := p[0]
			version := p[1]
			return name, version, true
		} else {
			fmt.Println("Error: invalid format")
		}
	}

	return "", "", false
}
