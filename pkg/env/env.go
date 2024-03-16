package env

import (
	"github.com/hit9/log"
	"github.com/joho/godotenv"
	"os"
)

var logger = log.Get("env")

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
