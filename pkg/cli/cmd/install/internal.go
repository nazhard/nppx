package install

import (
	"fmt"

	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/pkg/lockfile"
	"github.com/nazhard/nppx/pkg/resolver"
)

var ax string

func checkCache() (string, string, bool) {
	name, version, err := resolver.ReadDeps()
	if err != nil {
		fmt.Println(err)
	}

	ax = name + "_" + version
	cacheExist := fs.ReadDotModulesIsExist(ax)
	if cacheExist == true {
		return name, version, true
	}

	if cacheExist == false {
		name, version, err := resolver.ReadDevDeps()
		if err != nil {
			fmt.Println(err)
		}

		dx := name + "_" + version
		isCacheExist := fs.ReadDotModulesIsExist(dx)
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
			resolver.GetInfo(arg, "latest")
		}

		cacheExist := fs.ReadDotModulesIsExist(ax)
		if cacheExist == true {
			resolver.WriteDeps(name, version)
			return name, version, true
		}

		if cacheExist == false {
			name, version, err := resolver.ReadDevDeps()
			if err != nil {
				fmt.Println(err)
			}

			dx := name + "_" + version
			isCacheExist := fs.ReadDotModulesIsExist(dx)
			if isCacheExist {
				resolver.WriteDevDeps(name, version)
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
