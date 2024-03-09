package install

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/resolver"
	"github.com/nazhard/nppx/pkg/utils"
)

func Install(args []string) {
	if len(args) == 0 {
		err := utils.NoPackageJson()
		if err != nil {
			fmt.Println(err)
		}

		// modules.Search()
	}

	if len(args) != 0 {
		dl(args)

		/*
			m := modules.SearchArgs(args)
			if m == true {
				dl()
			}

			dd()
		*/
	}
}

func dl(args []string) {
	var wg sync.WaitGroup
	filePath := filepath.Join(setup.NPPX_PATH, ".modules.toml")
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			resolver.GetInfo(a)

			v := resolver.Version

			module_path := filepath.Join(setup.CACHE_PATH, a, v)
			fileName := module_path + "/" + a + "-" + v + ".tgz"

			_ = os.MkdirAll(module_path, os.ModePerm)
			resolver.Get(fileName, resolver.Tarball)

			_, err := file.WriteString(fmt.Sprintf("%s = \"v%s\"\n", a, v))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}(arg)
	}

	wg.Wait()
}
