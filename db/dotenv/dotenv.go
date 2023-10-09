package dotenv

import (
	"github.com/joho/godotenv"
)

func SetDotenv() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
