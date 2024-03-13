package install

import (
	"fmt"

	"github.com/nazhard/nppx/internal/fs"
	"github.com/nazhard/nppx/pkg/resolver"
)

func checkCache() (string, string, bool) {
	name, version, err := resolver.ReadPackageJson()
	if err != nil {
		fmt.Println(err)
	}

	ax := name + "_" + version
	cacheExist := fs.ReadDotModules(ax)
	if cacheExist == true {
		return name, version, true
	}

	return "", "", false
}

func installFromCache(a, v string) {
	fs.WriteToModulesJson(a, v)
}
