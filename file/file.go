// Package file
package file

import (
	"os"
	"path/filepath"
)

// ReadFile читает содержимое файла
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// IsJSONFile проверяет, имеет ли файл расширение .json.
func IsJSONFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".json"
}
