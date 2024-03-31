package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/nazhard/nppx/internal/http"
	"github.com/nazhard/nppx/internal/models"
)

var (
	Name    string
	Version string
	Tarball string
)

func PkgInfo(name, v string) {
	url := "https://registry.npmjs.org/" + name + "/" + v
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var info models.PkgInfo

	err = json.NewDecoder(data.Body).Decode(&info)
	if err != nil {
		fmt.Println(err)
	}

	Name = info.Name
	Version = info.Version
	Tarball = info.Dist.Tarball
}
