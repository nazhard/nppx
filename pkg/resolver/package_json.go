package resolver

import (
	"encoding/json"
	"fmt"
	"os"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func WriteDeps(newDeps, newDepsVersion string) {
	fileName := "package.json"
	existingData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var pkgJson PackageJSON
	err = json.Unmarshal(existingData, &pkgJson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Check if content already exists
	if existingVersion, exists := pkgJson.Dependencies[newDeps]; exists && existingVersion == newDepsVersion {
		fmt.Println("Warning: Content already exists.")
		return
	}
	// Append new content
	pkgJson.Dependencies[newDeps] = newDepsVersion

	// Marshal the updated data to JSON
	updatedData, err := json.MarshalIndent(pkgJson, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func WriteDevDeps(newDeps, newDepsVersion string) {
	fileName := "package.json"
	existingData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var pkgJson PackageJSON
	err = json.Unmarshal(existingData, &pkgJson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Check if content already exists
	if existingVersion, exists := pkgJson.DevDependencies[newDeps]; exists && existingVersion == newDepsVersion {
		fmt.Println("Warning: Content already exists.")
		return
	}
	// Append new content
	pkgJson.DevDependencies[newDeps] = newDepsVersion

	// Marshal the updated data to JSON
	updatedData, err := json.MarshalIndent(pkgJson, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(fileName, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content appended successfully.")
}
