package fs

import (
	"fmt"
	"path/filepath"

	"github.com/nazhard/nppx/internal/setup"
)

func WriteToDotModules(c string) {
	filePath := filepath.Join(setup.NPPX_Cache, ".modules.txt")

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

func ReadDotModules(c string) bool {
	filePath := filepath.Join(setup.NPPX_Cache, ".modules.txt")

	// Read existing content
	content, err := readExistingContent(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Check if new content already exists
	if contains(content, c) {
		return true
	}

	return false
}
