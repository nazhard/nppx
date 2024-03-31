package fs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nazhard/nppx/internal/setup"
)

func CreateSymlinks(jsonFile string) error {
	file, err := os.Open(jsonFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var symlinkMap map[string]string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&symlinkMap); err != nil {
		return err
	}

	for symlink, target := range symlinkMap {
		if _, err := os.Stat(setup.NPPX_Cache + target); os.IsNotExist(err) {
			return fmt.Errorf("Target path %s does not exist", target)
		}

		err := os.Symlink(setup.NPPX_Cache+target, ".nppx/modules/"+symlink)
		if err != nil {
			continue
		}

		//fmt.Printf("Symlink created: %s -> %s\n", symlink, target)
	}

	return nil
}

func WriteToModulesJson(c, cc string) {
	filePath := ".nppx/.modules.json"

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var existingData map[string]interface{}
	err = json.Unmarshal(content, &existingData)
	if err != nil {
		fmt.Println("Error unmarshaling existing content:", err)
		return
	}

	existingData[c] = "/" + c + "/" + cc

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(existingData)
	if err != nil {
		fmt.Println("Error encoding and writing to file:", err)
		return
	}
}

func WriteEmptyJSON() {
	data := make(map[string]interface{})

	file, err := os.Create(".nppx/.modules.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
