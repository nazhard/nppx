package install

import (
	"os"
	"sync"

	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/resolver"
)

func Install(args []string) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			resolver.GetLatestVersion(a)

			module_path := setup.CACHE_PATH + "/" + a + "/" + resolver.Version
			fileName := module_path + "/" + a + "-" + resolver.Version + ".tgz"

			_ = os.MkdirAll(module_path, os.ModePerm)
			resolver.Get(fileName, resolver.Tarball)
		}(arg)
	}

	wg.Wait()
}
