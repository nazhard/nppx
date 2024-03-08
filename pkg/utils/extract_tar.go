package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "file.tgz"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading tar header:", err)
			return
		}

		if header.Typeflag == tar.TypeReg {
			newFile, err := os.Create(header.Name)
			if err != nil {
				fmt.Println("Error creating new file:", err)
				return
			}
			defer newFile.Close()

			if _, err := io.Copy(newFile, tarReader); err != nil {
				fmt.Println("Error copying file content:", err)
				return
			}

			fmt.Println("Extracted:", header.Name)
		}
	}
	fmt.Println("Extraction complete.")
}
