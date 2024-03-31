package install

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/nazhard/nppx"
	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/internal/lockfile"
	"github.com/nazhard/nppx/internal/resolver"
	"github.com/nazhard/nppx/internal/setup"
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
	if b {
		resolver.PkgInfo(name, version)
	} else {
		resolver.PkgInfo(a, "latest")
	}
	return resolver.Name, resolver.Version
}

func ioStuff(name, version string) {
	modulePath := filepath.Join(setup.NPPX_Cache, name, version)
	nodeModules := filepath.Join(modulePath, "node_modules")
	fileName := filepath.Join(modulePath, fmt.Sprintf("%s-%s.tgz", name, version))

	err := os.MkdirAll(modulePath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	err = os.MkdirAll(nodeModules, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	nppx.Get(fileName, resolver.Tarball)
	fs.WriteToDotModules(fmt.Sprintf("%s_%s", name, version))
	nppx.ExtractGz(fileName, nodeModules, "package")

	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error removing file:", err)
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
			resolver.WriteDependencies(name, version, false)
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
			resolver.WriteDependencies(name, version, true)
			lockfile.WriteDeps(name, version, true)
		}(arg)
	}

	wg.Wait()
}

func cleanInstall() {
	var wg sync.WaitGroup

	n, v, err := resolver.ReadPackageJson()
	if err != nil {
		fmt.Println("Error reading dependencies:", err)
	}

	resolver.PkgInfo(n, v)

	ioStuff(n, v)

	lockfile.WriteDeps(n, v, false)

	fmt.Println(n, v)

	wg.Wait()
}
