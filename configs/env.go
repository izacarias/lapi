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
)

func GetConfigMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. %w", ErrMongoURI)
	}

	return os.Getenv("MONGOURI")
}

func GetConfigQueriesURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env file. %w", ErrResourceQueryURI)
		return "location/v3/queries"
	}
	return os.Getenv("RESOURCEURI")
}

func GetConfigQueriesZones() string {
	return GetConfigQueriesURI() + "/zones"
}
