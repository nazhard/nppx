package check

import (
	"errors"
	"os"
)

func NoPackageJson() error {
	_, err := os.Stat("package.json")
	current, _ := os.Getwd()

	if os.IsNotExist(err) {
		return errors.New("Error, no package.json found in " + current)
	}

	return nil
}

func CurrentDirNotEmpty() error {
	dirEntries, _ := os.ReadDir(".")

	if len(dirEntries) > 0 {
		return errors.New("Current working directory is not empty.")
	}

	return nil
}
