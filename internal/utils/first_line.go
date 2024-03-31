package utils

import (
	"bufio"
	"errors"
	"os"
)

func Check(path string) (string, error) {
	filePath := path

	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New(path + " doesn't exist")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return firstLine, nil
	} else {
		return "", errors.New("error reading file")
	}
}
