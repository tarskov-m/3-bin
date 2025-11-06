// Package api
package api

import (
	"fmt"

	"myapp/config"
)

type API struct {
	config *config.Config
}

func NewAPI(cfg *config.Config) *API {
	return &API{
		config: cfg,
	}
}

func (a *API) Start() {
	fmt.Printf("Starting API with key: %s\n", a.config.Key)
	// Используйте config здесь
}
