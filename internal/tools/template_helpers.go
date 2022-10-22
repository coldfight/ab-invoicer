package tools

import (
	"path/filepath"
)

func FullFilePath(relativePath string) string {
	abs, err := filepath.Abs(relativePath)
	if err != nil {
		return ""
	}
	return abs
}
