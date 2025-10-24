// Package file
package file

import (
	"os"
	"path/filepath"

	"myapp/interfaces"
)

// OSFileReader реализует интерфейс FileReader для работы с файловой системой
type OSFileReader struct{}

// NewOSFileReader создает новый экземпляр OSFileReader
func NewOSFileReader() interfaces.FileReader {
	return &OSFileReader{}
}

// ReadFile читает содержимое файла
func (f *OSFileReader) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// IsJSONFile проверяет, имеет ли файл расширение .json.
func (f *OSFileReader) IsJSONFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".json"
}

// Deprecated: используйте OSFileReader.ReadFile вместо этой функции
func ReadFile(path string) ([]byte, error) {
	reader := NewOSFileReader()
	return reader.ReadFile(path)
}

// Deprecated: используйте OSFileReader.IsJSONFile вместо этой функции
func IsJSONFile(path string) bool {
	reader := NewOSFileReader()
	return reader.IsJSONFile(path)
}
