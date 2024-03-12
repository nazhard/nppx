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

func dl(args []string, useDev bool) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			name, version, b := contains(a, "@", "/")

			if b == true {
				resolver.GetInfo(name, version)
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

			nppx.Extract(fileName, module_path, "package")
			err := os.Remove(fileName)
			if err != nil {
				fmt.Println(err)
			}

			if useDev == true {
				resolver.WriteDevDeps(a, v)
			} else {
				resolver.WriteDeps(a, v)
			}
		}(arg)
	}

	wg.Wait()
}
