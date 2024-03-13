package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Dist struct {
	Shasum  string `json:"shahum"`
	Tarball string `json:"tarball"`
}

type PkgInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Dist    Dist   `json:"dist"`
}

// package.json struct
type Pkg struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description,omitempty"`
	Main            string            `json:"main,omitempty"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Keywords        []string          `json:"keywords,omitempty"`
	Author          string            `json:"author,omitempty"`
	License         string            `json:"license,omitempty"`
}

var (
	Name    string
	Version string
	Tarball string
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func GetInfo(name, v string) {
	data, err := fetchData(name, v)
	if err != nil {
		fmt.Println("Error fetching data")
	}

	Name = data.Name
	Version = data.Version
	Tarball = data.Dist.Tarball
}

func fetchData(name, v string) (*PkgInfo, error) {
	var data PkgInfo

	url := "https://registry.npmjs.org/" + name + "/" + v

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("Error: " + name + " not found")
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
