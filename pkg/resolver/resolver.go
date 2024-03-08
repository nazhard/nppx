package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

var (
	Name        string
	Version     string
	Description string
	Tarball     string
)

func LatestVersion(name string) {
	data, err := fetchData(name)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	var vInfo VersionInfo

	lVersion := data.DistTags.Latest
	vInfo, exists := data.Versions[lVersion]

	if exists {
		Name = vInfo.Name
		Version = vInfo.Version
		Description = data.Description
		Tarball = vInfo.Dist.Tarball
	} else {
		fmt.Println("Latest version not found")
	}
}

func Get(name, url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file...")
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("Error copying content to file:", err)
		return
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
