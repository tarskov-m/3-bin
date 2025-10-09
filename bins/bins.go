package bins

import (
	"math/rand"
	"time"
)

type Bin struct {
	Id        string
	Name      string
	Private   bool
	CreatedAt time.Time
}

type BinList struct {
	Bins []Bin
}

func NewBin(name string, private bool) Bin {
	return Bin{
		Id:        generateID(),
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
