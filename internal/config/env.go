package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load();
	if err != nil {
		log.Println("Error while loading env file")
	}
}

func GetEnv(Key, fallback string) string {
	if value, ok := os.LookupEnv(Key); ok {
		return value
	}

	return fallback
}