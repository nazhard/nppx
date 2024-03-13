package install

import (
	"fmt"
	"os"
	"strings"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/lockfile"
	"github.com/urfave/cli/v2"
)

func Action(c *cli.Context) error {
	if c.Bool("D") == true {
		install(c.Args().Slice(), true)
	} else {
		install(c.Args().Slice(), false)
	}
	return nil
}

func install(args []string, useDev bool) {
	err := nppx.NoPackageJson()
	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		_, err := os.Stat(".nppx")
		if os.IsNotExist(err) {
			_ = os.Mkdir(".nppx", os.ModePerm)
			setup.DotNPPX()

			//fs.WriteEmptyJSON()
			lockfile.GenerateLockfile()
		}

		if len(args) == 0 {
			name, version, exist := checkCache()
			if exist == true {
				installFromCache(name, version)
				lockfile.WriteDeps(name, version, false)
			} else {
				fmt.Println("BAD")
			}
		}

		if len(args) != 0 {
			name, version, x := checkCacheWithArgs(args)
			if x == true {
				installFromCache(name, version)
				lockfile.WriteDeps(name, version, false)
			} else {
				fmt.Println("BAD")
				// dl(args, useDev)
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
