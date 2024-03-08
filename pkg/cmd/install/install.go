package install

import (
	"github.com/nazhard/nppx/pkg/resolver"
)

func Install(args string) {
	resolver.LatestVersion(args)

	resolver.Get(args+".tar.xz", resolver.Tarball)
}
