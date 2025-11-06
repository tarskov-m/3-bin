// Package interfaces содержит определения интерфейсов для dependency injection
package interfaces

import "myapp/models"

// Storage интерфейс для работы с хранилищем данных
type Storage interface {
	SaveBin(bin models.Bin) error
	LoadBins() ([]models.Bin, error)
}

// FileReader интерфейс для чтения файлов
type FileReader interface {
	ReadFile(path string) ([]byte, error)
	IsJSONFile(path string) bool
}

// BinService интерфейс для работы с bin объектами
type BinService interface {
	CreateBin(name string, private bool) models.Bin
	GenerateID() string
}
