// Package storage
package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"myapp/bins"
)

const fileName = "bins.json"

func SaveBin(bin bins.Bin) error {
	binsList, err := LoadBins()
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

func LoadBins() ([]bins.Bin, error) {
	data, err := os.ReadFile(filepath.Join(".", fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return []bins.Bin{}, nil
		}
		return nil, err
	}
	var binsList []bins.Bin
	err = json.Unmarshal(data, &binsList)
	return binsList, err
}
