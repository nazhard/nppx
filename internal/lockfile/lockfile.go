package lockfile

import (
	"encoding/json"
	"fmt"
	"os"
)

type Lockfile struct {
	Dependencies map[string]Dependency `json:"dependencies"`
	Packages     map[string]Package    `json:"packages"`
}

type Dependency struct {
	Specifier string `json:"specifier"`
	Version   string `json:"version"`
	Module    string `json:"module"`
	Dev       bool   `json:"dev"`
}

type Package struct {
	Integrity            string            `json:"integrity"`
	HasBin               bool              `json:"hasBin"`
	RequiresBuild        string            `json:"requiresBuild"`
	Dependencies         map[string]string `json:"dependencies"`
	OptionalDependencies map[string]string `json:"optionalDependencies"`
	Dev                  bool              `json:"dev"`
}

func WriteDeps(newDeps, newDepsVersion string, isDev bool) {
	file := "nppx-lock.json"
	existingData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var lockfile Lockfile
	err = json.Unmarshal(existingData, &lockfile)
	if err != nil {
		fmt.Println("Error unmarshaling nppx-lock.json", err)
		return
	}

	// Check if content already exists
	var dependency Dependency
	if existingVersion, exists := lockfile.Dependencies[newDeps]; exists && existingVersion.Version == dependency.Version {
		fmt.Println("Content already go")
		return
	}
	// Append new content
	lockfile.Dependencies[newDeps] = Dependency{
		Version: newDepsVersion,
		Module:  "/" + newDeps + "/" + newDepsVersion + "/",
		Dev:     isDev,
	}

	// Marshal the updated data to JSON
	updatedData, err := json.MarshalIndent(lockfile, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(file, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
