package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

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

// WriteDependencies writes dependencies to the package.json file.
func WriteDependencies(newDeps, newDepsVersion string, dev bool) error {
	fileName := "package.json"
	existingData, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var pkgJson Pkg
	err = json.Unmarshal(existingData, &pkgJson)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	var deps map[string]string
	if dev {
		deps = pkgJson.DevDependencies
		pkgJson.MergeDevDependencies()
	} else {
		deps = pkgJson.Dependencies
		pkgJson.MergeDependencies()
	}

	// Check if content already exists
	if existingVersion, exists := deps[newDeps]; exists && existingVersion == newDepsVersion {
		return errors.New("content already exists")
	}

	deps[newDeps] = newDepsVersion

	updatedData, err := json.MarshalIndent(pkgJson, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling package.json: %v", err)
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(fileName, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

// ReadPackageJson reads package.json and returns a dependency and its version.
func ReadPackageJson() (string, string, error) {
	file := "package.json"
	data, err := os.ReadFile(file)
	if err != nil {
		return "", "", fmt.Errorf("error reading package.json: %v", err)
	}

	var pkg Pkg
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		return "", "", fmt.Errorf("error unmarshaling package.json: %v", err)
	}

	// Just returning the first dependency found for demonstration
	for dep, version := range pkg.Dependencies {
		return dep, version, nil
	}
	return "", "", errors.New("no dependencies found")
}

// MergeDependencies  merges dependencies from devDependencies into dependencies.
func (d *Pkg) MergeDependencies() {
	for key, value := range d.DevDependencies {
		if _, exists := d.Dependencies[key]; !exists {
			d.Dependencies[key] = value
			delete(d.DevDependencies, key)
		}
	}
}

// Check for duplicate modules, then delete them from the `devDependencies` section
func (d *Pkg) MergeDevDependencies() {
	for key, value := range d.Dependencies {
		if _, ok := d.Dependencies[key]; ok {
			d.DevDependencies[key] = value
			delete(d.Dependencies, key)
		}
	}
}
