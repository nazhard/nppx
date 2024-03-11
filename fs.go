package nppx

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nazhard/nppx/internal/setup"
)

func WriteToDotModules(c string) {
	filePath := filepath.Join(setup.NPPX_PATH, ".modules.txt")

	// Read existing content
	existingContent, err := readExistingContent(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Check if new content already exists
	if contains(existingContent, c) {
		fmt.Println("Warning: Content already exists.")
		return
	}

	// Append new content
	if err := appendToFile(filePath, c); err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Content added successfully.")
}

func readExistingContent(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func appendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	return err
}

func CreateSymlinks(jsonFile string) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var symlinkMap map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&symlinkMap); err != nil {
		return err
	}

	for symlink, target := range symlinkMap {
		if _, err := os.Stat(setup.CACHE_PATH + target); os.IsNotExist(err) {
			return fmt.Errorf("Target path %s does not exist", target)
		}

		err := os.Symlink(setup.CACHE_PATH+target, symlink)
		if err != nil {
			return err
		}

		fmt.Printf("Symlink created: %s -> %s\n", symlink, target)
	}

	return nil
}

//func main() {
//	jsonFile := "sym.json"
//	err := createSymlinks(jsonFile)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//}
