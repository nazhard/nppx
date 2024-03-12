package setup

import (
	"os"
	"sync"
)

func DotNPPX() {
	dir1 := ".nppx/modules"
	dirNames := []string{dir1}

	f1 := ".nppx/.modules.json"
	fileNames := []string{f1}

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
}

func CreateDotNPPX() error {
	err := os.MkdirAll(".nppx", os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
