package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

var (
	Name    string
	Version string
	Tarball string
)

var client = &http.Client{
	Timeout: 3 * time.Second,
}

func GetLatestVersion(name string) {
	resp, err := fetchLatestVersion(name)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	data, err := fetchData(name, resp.DistTags.Latest)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}

	Name = data.Name
	Version = data.Version
	Tarball = data.Dist.Tarball
}

func Get(name, url string) {
	res, err := client.Get(url)
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

func fetchLatestVersion(name string) (*Pkg, error) {
	var data Pkg

	url := "https://registry.npmjs.org/" + name

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func fetchData(name, v string) (*VersionInfo, error) {
	var data VersionInfo

	url := "https://registry.npmjs.org/" + name + "/" + v

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
