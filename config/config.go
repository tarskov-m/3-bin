// Package config читает env файл
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Ошибка: .env файл не найден: %v\n", err)
	}

	return &Config{
		Key: os.Getenv("KEY"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
