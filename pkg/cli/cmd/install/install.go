package install

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/internal/setup"
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

	t := time.Since(time.Now())

	if err == nil {
		_, err := os.Stat(".nppx")
		if os.IsNotExist(err) {
			_ = os.Mkdir(".nppx", os.ModePerm)
			setup.DotNPPX()

			fs.WriteEmptyJSON()
		}

		if len(args) == 0 {
			name, version, exist := checkCache()
			if exist == true {
				installFromCache(name, version)
				fmt.Printf("Done in %s, Installed from cache \n", t)
			} else {
				fmt.Println("BAD")
			}
		}

		if len(args) != 0 {
			name, version, x := checkCache()
			if x == true {
				installFromCache(name, version)
				fmt.Printf("Done in %s, Installed from cache \n", t)
			} else {
				fmt.Println("BAD")
				// dl(args, useDev)
			}
		}

		err = fs.CreateSymlinks(".nppx/.modules.json")
		if err != nil {
			fmt.Println("Error:", err)
		}
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
