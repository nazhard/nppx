package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

	var pkgJson Pkg
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

	pkgJson.Dependencies[newDeps] = newDepsVersion

	pkgJson.MergeDependencies()

	updatedData, err := json.MarshalIndent(pkgJson, "", "  ")

	fmt.Println(string(updatedData))
	if err != nil {
		fmt.Println("Error marshaling package.json:", err)
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

	var pkgJson Pkg
	err = json.Unmarshal(existingData, &pkgJson)
	if err != nil {
		fmt.Println("Error unmarshaling package.json:", err)
		return
	}

	// Check if content already exists
	if existingVersion, exists := pkgJson.DevDependencies[newDeps]; exists && existingVersion == newDepsVersion {
		fmt.Println("Warning: Content already exists.")
		return
	}
	// Append new content
	pkgJson.DevDependencies[newDeps] = newDepsVersion

	pkgJson.MergeDevDependencies()

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

// This function reads each dependencies.
// TODO: make it can reads dev, peer, and optional dependencies.
func ReadPackageJson() (string, string, error) {
	file := "package.json"
	data, err := os.ReadFile(file)
	if err != nil {
		return "", "", errors.New("Error reading package.json")
	}

	var pkg Pkg
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		return "", "", errors.New("Error reading package.nson:")
	}

	var (
		a string
		b string
	)

	for i, c := range pkg.Dependencies {
		a, b = i, c
	}

	return a, b, nil
}

// Check for duplicate modules, then delete them from the `dependencies` section
func (d *Pkg) MergeDependencies() {
	for key, value := range d.DevDependencies {
		if _, ok := d.DevDependencies[key]; ok {
			delete(d.DevDependencies, key)
			d.Dependencies[key] = value
		}
	}
}

// Check for duplicate modules, then delete them from the `devDependencies` section
func (d *Pkg) MergeDevDependencies() {
	for key, value := range d.Dependencies {
		if _, ok := d.Dependencies[key]; ok {
			delete(d.Dependencies, key)
			d.DevDependencies[key] = value
		}
	}
}
