// Package storage
package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"myapp/bins"
	"myapp/interfaces"
	"myapp/models"
)

const fileName = "bins.json"

// JSONStorage реализует интерфейс Storage для работы с JSON файлами
type JSONStorage struct{}

// NewJSONStorage создает новый экземпляр JSONStorage
func NewJSONStorage() interfaces.Storage {
	return &JSONStorage{}
}

func (s *JSONStorage) SaveBin(bin models.Bin) error {
	binsList, err := s.LoadBins()
	if err != nil {
		return err
	}
	binsList = append(binsList, bin)
	data, err := json.MarshalIndent(binsList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(".", fileName), data, 0644)
}

func (s *JSONStorage) LoadBins() ([]models.Bin, error) {
	data, err := os.ReadFile(filepath.Join(".", fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Bin{}, nil
		}
		return nil, err
	}
	var binsList []models.Bin
	err = json.Unmarshal(data, &binsList)
	return binsList, err
}

// Deprecated: используйте JSONStorage.SaveBin вместо этой функции
func SaveBin(bin bins.Bin) error {
	storage := NewJSONStorage()
	// Конвертируем bins.Bin в models.Bin
	modelBin := models.Bin{
		ID:        bin.ID,
		Name:      bin.Name,
		Private:   bin.Private,
		CreatedAt: bin.CreatedAt,
	}
	return storage.SaveBin(modelBin)
}

// Deprecated: используйте JSONStorage.LoadBins вместо этой функции
func LoadBins() ([]bins.Bin, error) {
	storage := NewJSONStorage()
	modelBins, err := storage.LoadBins()
	if err != nil {
		return nil, err
	}
	// Конвертируем []models.Bin в []bins.Bin
	var binsList []bins.Bin
	for _, modelBin := range modelBins {
		binsList = append(binsList, bins.Bin{
			ID:        modelBin.ID,
			Name:      modelBin.Name,
			Private:   modelBin.Private,
			CreatedAt: modelBin.CreatedAt,
		})
	}
	return binsList, nil
}
