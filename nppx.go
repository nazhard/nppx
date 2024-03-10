package nppx

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(name, url string) {
	res, err := http.Get(url)
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
