package resolver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VersionInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Pkg struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Versions    map[string]VersionInfo `json:"versions"`
	DistTags    struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

func LatestVersion() {
	data, err := fetchData()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	fmt.Printf("Latest Version: %s, Name: %s\n", data.DistTags.Latest, data.Name)
}

func AllVersion() {
	data, err := fetchData()
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	fmt.Println("Versions:")
	for version, info := range data.Versions {
		fmt.Printf("  Version: %s, Name: %s\n", version, info.Name)
	}
}

func fetchData() (*Pkg, error) {
	url := "https://registry.npmjs.org/express"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data Pkg
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
