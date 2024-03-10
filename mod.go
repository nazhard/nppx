package nppx

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nazhard/nppx/internal/setup"
)

var line string

func ReadDotModules(name string) string {
	file, err := os.Open(setup.NPPX_PATH + "/.modules.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, name) {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return line
}
