package nppx

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/codeclysm/extract/v3"
)

func Extract(in, out, path string) {
	file, _ := os.Open(in)
	var shift = func(path string) string {
		parts := strings.Split(path, string(filepath.Separator))
		parts = parts[1:]
		return strings.Join(parts, string(filepath.Separator))
	}
	extract.Gz(context.TODO(), file, out, shift)
}
