// Package bins реализует функциональность для создания и управления
// объектами Bin, представляющими собой контейнеры с уникальными
// идентификаторами и возможностью установки приватности.
package bins

import (
	"math/rand"
	"time"
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

func NewBin(name string, private bool) Bin {
	return Bin{
		ID:        generateID(),
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const idLength = 8

	b := make([]byte, idLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
