package resolver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dist struct {
	Tarball string `json:"tarball"`
}

type VersionInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Dist    Dist   `json:"dist"`
}

type Pkg struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Versions    map[string]VersionInfo `json:"versions"`
	DistTags    struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

func LatestVersion(name string) {
	data, err := fetchData(name)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	var vInfo VersionInfo

	lVersion := data.DistTags.Latest
	vInfo, exists := data.Versions[lVersion]

	if exists {
		pName := vInfo.Name
		pVersion := vInfo.Version

		fmt.Printf("Latest Version: %s, Name: %s\n", pVersion, pName)
	} else {
		fmt.Println("Latest version not found")
	}
}

func fetchData(module string) (*Pkg, error) {
	url := "https://registry.npmjs.org/" + module

	var data Pkg

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
