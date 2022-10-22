package tools

import (
	"fmt"
	"path/filepath"
)

func FullFilePath(relativePath string) string {
	abs, err := filepath.Abs(relativePath)
	if err != nil {
		return ""
	}
	return abs
}

func Currency(num float64) string {
	return fmt.Sprintf("$%.2f", num)
}
