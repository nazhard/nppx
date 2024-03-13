package nppx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/codeclysm/extract/v3"
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

func ExtractGz(in, out, path string) {
	file, _ := os.Open(in)
	var shift = func(path string) string {
		parts := strings.Split(path, string(filepath.Separator))
		parts = parts[1:]
		return strings.Join(parts, string(filepath.Separator))
	}
	extract.Gz(context.TODO(), file, out, shift)
}

func ExtractXz(in, out string) {
	file, _ := os.Open(in)
	extract.Xz(context.TODO(), file, out, nil)
}
