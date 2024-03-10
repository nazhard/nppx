package nppx

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
)

func ExtractTarXZ(tarFile, extractPath string) error {
	file, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer file.Close()

	xzReader, err := xz.NewReader(file)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(xzReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		targetPath := filepath.Join(extractPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// Create directory if it doesn't exist
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			// Extract regular file
			file, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}

func extractTarGzip(tgzFile, extractPath string) error {
	file, err := os.Open(tgzFile)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		targetPath := filepath.Join(extractPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// Create directory if it doesn't exist
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			// Extract regular file
			file, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}

/*
func main() {
	tgzFile := "file.tgz"
	extractPath := "extracted/dir"

	err := extractDir(tgzFile, extractPath)
	if err != nil {
		fmt.Println("Error extracting:", err)
	} else {
		fmt.Println("Extraction successful!")
	}

	tarFile := "file.tar.xz"
	extractPath := "extracted/dir"

	err := extractDir(tarFile, extractPath)
	if err != nil {
		fmt.Println("Error extracting:", err)
	} else {
		fmt.Println("Extraction successful!")
	}
}
*/
