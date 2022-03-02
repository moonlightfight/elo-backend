package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDbUri() string {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	mongodbUri := os.Getenv("MONGODB_URI")

	return mongodbUri
}

func GetDbName() string {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	databaseName := os.Getenv("DATABASE_NAME")

	return databaseName
}
