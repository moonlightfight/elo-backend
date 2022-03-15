package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDbUri() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	mongodbUri := os.Getenv("MONGODB_URI")

	return mongodbUri
}

func GetEnvVar(varName string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	key := os.Getenv(varName)

	return key
}
