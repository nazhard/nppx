package lockfile

import (
	"encoding/json"
	"fmt"
	"os"
)

// This function will be generates lockfile and write
// first data to lockfile
func genLockfile() {
	lockfile := Lockfile{
		Dependencies: make(map[string]Dependency),
		Packages:     make(map[string]Package),
	}

	data, err := json.MarshalIndent(lockfile, "", "  ")
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	file, err := os.Create("nppx-lock.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// This function checks if lockfile is exist or not.
// If not, the lockfile will be generate by generateLockfile()
func GenerateLockfile() {
	_, err := os.Stat("nppx-lock.json")
	if os.IsNotExist(err) {
		genLockfile()
	}
}
