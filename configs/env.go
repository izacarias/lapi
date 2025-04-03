package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrMongoURI         = errors.New("MONGOURI not found in .env")
	ErrResourceQueryURI = errors.New("RESOURCEURI not found in .env")
	ErrApiVersion       = errors.New("APIVERSION not found in .env")
	ErrPurgeDatabase    = errors.New("PURGEDATABASE not found in .env")
	ErrInsertMockData   = errors.New("INSERTMOCKDATA not found in .env")
)

func GetConfigMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrMongoURI)
	}

	return os.Getenv("MONGOURI")
}

func GetConfigApiRoot() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrResourceQueryURI)
	}
	return os.Getenv("APIROOT")
}

func GetConfigApiVersion() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrResourceQueryURI)
	}
	return os.Getenv("APIVERSION")
}

func GetConfigClearDatabase() bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrPurgeDatabase)
	}

	purgeDB := os.Getenv("PURGEDATABASE")
	return purgeDB == "1"
}

func GetConfigInsertMockData() bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrInsertMockData)
	}

	insertMockData := os.Getenv("INSERTMOCKDATA")
	return insertMockData == "1"
}
