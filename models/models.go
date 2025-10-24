// Package models содержит основные модели данных
package models

import "time"

// Bin представляет контейнер с уникальным идентификатором
type Bin struct {
	ID        string
	Name      string
	Private   bool
	CreatedAt time.Time
}

// BinList представляет список bins
type BinList struct {
	Bins []Bin
}
