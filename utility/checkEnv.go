package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CheckEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv(name)
	return env
}
