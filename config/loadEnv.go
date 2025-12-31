package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv akan mencoba load file .env jika ada.
// Tidak fatal jika file tidak ditemukan.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, skipping...")
	}
}

// GetEnv mengambil value dari environment variable.
// Fatal jika variable tidak ditemukan.
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ ENV %s is missing", key)
	}
	return value
}
