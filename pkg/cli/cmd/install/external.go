package install

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/lockfile"
	"github.com/nazhard/nppx/pkg/resolver"
)

func do(args []string, action string, have bool) {
	switch have {
	case true:
		switch action {
		case "dev":
			getDev(args)
		case "def":
			getDef(args)
		}
	case false:
		switch action {
		case "dev":
			cleanDev()
		case "def":
			cleanDef()
		}
	}
}

func fetch(a string) (string, string) {
	name, version, b := contains(a, "@", "/")
	if b == true {
		resolver.GetInfo(name, version)
		return name, version
	} else {
		resolver.GetInfo(a, "latest")
		name, version := resolver.Name, resolver.Version
		return name, version
	}
}

func ioStuff(name, version string) {
	module_path := filepath.Join(setup.CACHE_PATH, name, version)
	fileName := module_path + "/" + name + "-" + version + ".tgz"

	_ = os.MkdirAll(module_path, os.ModePerm)
	nppx.Get(fileName, resolver.Tarball)

	fs.WriteToDotModules(fmt.Sprintf("%s_%s", name, version))

	_ = nppx.ReadDotModules(name)

	nppx.ExtractGz(fileName, module_path, "package")
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func getDef(args []string) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			name, version := fetch(a)
			ioStuff(name, version)
			lockfile.WriteDeps(name, version, false)
		}(arg)
	}

	wg.Wait()
}

func getDev(args []string) {
	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		go func(a string) {
			defer wg.Done()

			name, version := fetch(a)
			ioStuff(name, version)
			lockfile.WriteDeps(name, version, true)
		}(arg)
	}

	wg.Wait()
}

func cleanDef() {}
func cleanDev() {}
