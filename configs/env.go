package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// It loads the environment variables from the .env file and returns the value of the MONGO_URI
// variable
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment")
	}
	log.Println("Load environment successfully")
	return os.Getenv("MONGO_URI")
}
func EnvJwtKey() string {
	return os.Getenv("JWT_KEY")
}
