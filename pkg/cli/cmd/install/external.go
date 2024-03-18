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
	if have {
		switch action {
		case "dev":
			getDev(args)
		case "def":
			getDef(args)
		}
	} else {
		cleanInstall()
	}
}

func fetch(a string) (string, string) {
	name, version, b := contains(a, "@", "/")
	if b == true {
		resolver.GetInfo(name, version)
		name, version := resolver.Name, resolver.Version
		return name, version
	} else {
		resolver.GetInfo(a, "latest")
		name, version := resolver.Name, resolver.Version
		return name, version
	}
}

func ioStuff(name, version string) {
	module_path := filepath.Join(setup.CACHE_PATH, name, version)
	node_modules := filepath.Join(module_path, "node_modules")
	fileName := module_path + "/" + name + "-" + version + ".tgz"

	_ = os.MkdirAll(module_path, os.ModePerm)
	_ = os.MkdirAll(node_modules, os.ModePerm)
	nppx.Get(fileName, resolver.Tarball)

	fs.WriteToDotModules(fmt.Sprintf("%s_%s", name, version))

	// _ = nppx.ReadDotModules(name)

	nppx.ExtractGz(fileName, node_modules, "package")
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
			resolver.WriteDeps(name, version)
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
			resolver.WriteDevDeps(name, version)
			lockfile.WriteDeps(name, version, true)
		}(arg)
	}

	wg.Wait()
}

func cleanInstall() {
	var wg sync.WaitGroup

	n, v, err := resolver.ReadDeps()
	if err != nil {
		fmt.Println(err)
	}
	dN, dV, err := resolver.ReadDevDeps()
	if err != nil {
		fmt.Println(err)
	}

	resolver.GetInfo(n, v)
	resolver.GetInfo(dN, dV)

	ioStuff(n, v)
	ioStuff(dN, dV)

	lockfile.WriteDeps(n, v, false)
	lockfile.WriteDeps(dN, dV, true)

	fmt.Println(n, v, dN, dV)

	wg.Wait()
}
