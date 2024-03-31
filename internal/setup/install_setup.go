package setup

import (
	"os"
)

func CreateDotNPPXDir() error {
	if err := os.MkdirAll(".nppx", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(".nppx/modules", os.ModePerm); err != nil {
		return err
	}

	return nil
}
