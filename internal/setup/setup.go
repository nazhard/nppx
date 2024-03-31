package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	Home, _      = os.UserHomeDir()
	NPPX_Cache   = filepath.Join(Home, ".nppx-caches")
	NPPX_Modules = filepath.Join(Home, ".nppx-modules")
	NPPX_Bin     = filepath.Join(NPPX_Modules, "bin")
)

func IsNeeded() bool {
	_, err := os.Stat(NPPX_Modules)
	_, errr := os.Stat(NPPX_Cache)
	if os.IsNotExist(err) && os.IsNotExist(errr) {
		_ = os.MkdirAll(NPPX_Modules, os.ModePerm)
		_ = os.MkdirAll(NPPX_Cache, os.ModePerm)

		runner()

		return false
	}
	return true
}

func runner() {
	fmt.Println("Wait for initialization...")

	nppxInit()

	fmt.Println("Try type \"nppx help\"")
}

func nppxInit() {
	dirNames := []string{NPPX_Modules, NPPX_Cache, NPPX_Bin}
	fileNames := []string{filepath.Join(NPPX_Cache, ".modules.txt")}

	for _, dirName := range dirNames {
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dirName, err)
			return
		}
	}

	for _, fileName := range fileNames {
		if _, err := os.Create(fileName); err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			return
		}
	}

	fmt.Println("All files and directories created successfully!")
}
