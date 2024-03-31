package install

import (
	"fmt"

	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/internal/lockfile"
	"github.com/nazhard/nppx/internal/resolver"
)

var ax string

func checkCache() (string, string, bool) {
	name, version, err := resolver.ReadPackageJson()
	if err != nil {
		fmt.Println(err)
	}

	ax = name + "_" + version
	cacheExist := fs.ReadDotModules(ax)
	if cacheExist == true {
		return name, version, true
	}

	if cacheExist == false {
		name, version, err := resolver.ReadPackageJson()
		if err != nil {
			fmt.Println(err)
		}

		dx := name + "_" + version
		isCacheExist := fs.ReadDotModules(dx)
		if isCacheExist == true {
			return name, version, true
		}
	}

	return "", "", false
}

func checkCacheWithArgs(args []string) (string, string, bool) {
	for _, arg := range args {
		name, version, b := contains(arg, "@", "/")

		if b == true {
			ax = name + "_" + version
		} else {
			resolver.PkgInfo(arg, "latest")
		}

		cacheExist := fs.ReadDotModules(ax)
		if cacheExist == true {
			resolver.WriteDependencies(name, version, false)
			return name, version, true
		}

		if cacheExist == false {
			name, version, err := resolver.ReadPackageJson()
			if err != nil {
				fmt.Println(err)
			}

			dx := name + "_" + version
			isCacheExist := fs.ReadDotModules(dx)
			if isCacheExist {
				resolver.WriteDependencies(name, version, true)
				return name, version, true
			}
		}
	}
	return "", "", false
}

func installFromCache(n, v, action string) {
	switch action {
	case "dev":
		go installCache(n, v, true)
	case "def":
		go installCache(n, v, false)
	}
}

func installCache(name, version string, isDev bool) {
	fs.CreateSymlinks("nppx-lock.json")
	lockfile.WriteDeps(name, version, isDev)
}
