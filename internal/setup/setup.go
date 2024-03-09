package setup

import (
	"fmt"
	"os"
	"sync"
)

var (
	PATH       = os.Getenv("NPPX_HOME")
	NPPX_PATH  = PATH + "/.nppx"
	CACHE_PATH = NPPX_PATH + "/.cache"
	BIN_PATH   = PATH + "/bin"
)

func InitFunc() bool {
	_, err := os.Stat(NPPX_PATH)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(NPPX_PATH, os.ModePerm)

		runner()

		return false
	}
	return true
}

func runner() {
	fmt.Println("Wait for initialisation...")

	nppxInit()

	fmt.Println("Initialize completed.")
	fmt.Println("Try type \"nppx help\"")
}

func nppxInit() {
	dirNames := []string{BIN_PATH, CACHE_PATH}

	xv := NPPX_PATH + "/.modules.toml"
	fileNames := []string{xv}

	var wg sync.WaitGroup

	for _, dirName := range dirNames {
		wg.Add(1)
		go createDirectory(dirName, &wg)
	}

	for _, fileName := range fileNames {
		wg.Add(1)
		go createFile(fileName, &wg)
	}

	wg.Wait()

	fmt.Println("All files and directories created successfully!")
}

func createFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	fmt.Printf("File %s created successfully!\n", fileName)
}

func createDirectory(dirName string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dirName, err)
		return
	}

	fmt.Printf("Directory %s created successfully!\n", dirName)
}
