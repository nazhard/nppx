package resolver

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nazhard/nppx/pkg/cli/cmd/initial"
)

// This function writes dependencies when user does `nppx i <package>`
// to the package.json file in the `dependencies` section
func WriteDeps(newDeps, newDepsVersion string) {
	fileName := "package.json"
	existingData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var pkgJson initial.Pkg
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

// This function writes dependencies when user does `nppx i -D <package>` (using -D flag)
// to the package.json file in the `devDependencies` section
func WriteDevDeps(newDeps, newDepsVersion string) {
	fileName := "package.json"
	existingData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var pkgJson initial.Pkg
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
