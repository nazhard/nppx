package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

var (
	Name    string
	Version string
	Tarball string
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func GetInfo(name string) {
	data, err := fetchData(name, "latest")
	if err != nil {
		fmt.Println("Error fetching data")
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

	if res.StatusCode == 200 {
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
	} else if res.StatusCode == 404 {
		fmt.Printf("Error: %s not found", name)
		return
	} else {
		fmt.Printf("Error when fetching %s", name)
		return
	}
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
