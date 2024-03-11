package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/resolver"
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

			nppx.WriteEmptyJSON()
		}
	}

	if len(args) == 0 {
		// modules.Search()
	}

	if len(args) != 0 {
		dl(args, useDev)

		/*
			m := modules.SearchArgs(args)
			if m == true {
				dl()
			}

			dd()
		*/
	}

	err = nppx.CreateSymlinks(".nppx/.modules.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func dl(args []string, useDev bool) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			if strings.Contains(a, "@") && !strings.Contains(a, "/") {
				p := strings.Split(a, "@")
				if len(p) == 2 {
					name := p[0]
					version := p[1]
					resolver.GetInfo(name, version)
					fmt.Println("version: " + name + version)
				} else {
					fmt.Println("Invalid format")
				}
			} else {
				resolver.GetInfo(a, "latest")
			}

			v := resolver.Version

			module_path := filepath.Join(setup.CACHE_PATH, a, v)
			fileName := module_path + "/" + a + "-" + v + ".tgz"

			_ = os.MkdirAll(module_path, os.ModePerm)
			nppx.Get(fileName, resolver.Tarball)

			nppx.WriteToDotModules(fmt.Sprintf("%s_%s", a, v))

			_ = nppx.ReadDotModules(a)
			nppx.WriteToModulesJson(a, v)

			if useDev == true {
				resolver.WriteDevDeps(a, v)
			} else {
				resolver.WriteDeps(a, v)
			}
		}(arg)
	}

	wg.Wait()
}
