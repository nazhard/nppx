package install

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/resolver"
)

func Install(args []string, useDev bool) {
	err := nppx.NoPackageJson()
	if err != nil {
		fmt.Println(err)
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
}

func dl(args []string, useDev bool) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			resolver.GetInfo(a)

			v := resolver.Version

			module_path := filepath.Join(setup.CACHE_PATH, a, v)
			//fileName := module_path + "/" + a + "-" + v + ".tgz"

			_ = os.MkdirAll(module_path, os.ModePerm)
			//nppx.Get(fileName, resolver.Tarball)

			nppx.WriteToDotModules(fmt.Sprintf("%s_%s", a, v))

			l := nppx.ReadDotModules(a)
			fmt.Println(l)

			if useDev == true {
				resolver.WriteDevDeps(a, v)
			} else {
				resolver.WriteDeps(a, v)
			}
		}(arg)
	}

	wg.Wait()
}
