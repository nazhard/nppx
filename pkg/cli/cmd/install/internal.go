package install

import (
	"fmt"

	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/pkg/resolver"
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
		cacheExist := fs.ReadDotModules(ax)
		if cacheExist == true {
			resolver.WriteDeps(name, version)
			return name, version, true
		}
	}
	return "", "", false
}

func installFromCache(a, v string) {
	fs.WriteToModulesJson(a, v)
}
