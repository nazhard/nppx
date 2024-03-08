package install

import (
	"os"

	"github.com/nazhard/nppx/internal/setup"
	"github.com/nazhard/nppx/pkg/resolver"
)

func Install(args string) {
	resolver.GetLatestVersion(args)

	module_path := setup.NPPX_PATH + "/" + args + "/" + resolver.Version
	fileName := module_path + "/" + args + "-" + resolver.Version + ".tgz"

	_ = os.MkdirAll(module_path, os.ModePerm)
	resolver.Get(fileName, resolver.Tarball)
}
