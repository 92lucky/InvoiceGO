package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
}

	func GetEnv(key string) string {
		value := os.Getenv(key)
		if value ==""{
			log.Fatalf("ENV %s is running", key)
		}
		return value
	}


