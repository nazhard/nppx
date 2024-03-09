package utils

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
