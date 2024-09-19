package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	return os.Getenv(key)
}