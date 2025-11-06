// Package bins реализует функциональность для создания и управления
// объектами Bin, представляющими собой контейнеры с уникальными
// идентификаторами и возможностью установки приватности.
package bins

import (
	"math/rand"
	"time"

	"myapp/interfaces"
	"myapp/models"
)

type Bin struct {
	ID        string
	Name      string
	Private   bool
	CreatedAt time.Time
}

type BinList struct {
	Bins []Bin
}

// BinServiceImpl реализует интерфейс BinService
type BinServiceImpl struct{}

// NewBinService создает новый экземпляр BinService
func NewBinService() interfaces.BinService {
	return &BinServiceImpl{}
}

func (bs *BinServiceImpl) CreateBin(name string, private bool) models.Bin {
	return models.Bin{
		ID:        bs.GenerateID(),
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func (bs *BinServiceImpl) GenerateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const idLength = 8

	b := make([]byte, idLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Deprecated: используйте BinService.CreateBin вместо этой функции
func NewBin(name string, private bool) Bin {
	service := NewBinService()
	modelBin := service.CreateBin(name, private)
	// Конвертируем models.Bin в bins.Bin
	return Bin{
		ID:        modelBin.ID,
		Name:      modelBin.Name,
		Private:   modelBin.Private,
		CreatedAt: modelBin.CreatedAt,
	}
}

// Deprecated: используйте BinService.GenerateID вместо этой функции
func generateID() string {
	service := NewBinService()
	return service.GenerateID()
}
