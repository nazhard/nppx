package setup

import (
	"fmt"
	"os"
)

var (
	PATH      = os.Getenv("NPPX_HOME")
	NPPX_PATH = PATH + "/node_modules/.nppx"
	BIN_PATH  = PATH + "/bin"
)

func InitFunc() bool {
	_, err := os.Stat(NPPX_PATH)
	if os.IsNotExist(err) {
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
	_, err := os.Stat(NPPX_PATH)
	if os.IsNotExist(err) {
		fmt.Println("Adding nppx path...")
		_ = os.MkdirAll(NPPX_PATH, os.ModePerm)

		fmt.Println("Adding nppx files...")
		a, _ := os.Create(NPPX_PATH + "/.modules.toml")
		defer a.Close()
	}
}
