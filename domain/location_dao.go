package domain

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var locationCollection *mongo.Collection = configs.GetCollection(configs.DB, "locations")

type LocationMongo struct {
	ElementType string  `bson:"element_type"`
	ElementId   string  `bson:"element_id"`
	Latitude    float32 `bson:"latitude"`
	Longitude   float32 `bson:"longitude"`
	Altitude    float32 `bson:"altitude"`
	Timestamp   int64   `bson:"timestamp"`
}

func GetLocation(elementType string, elementId string) (*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"element_type": elementType, "element_id": elementId}
	cursor, err := locationCollection.Find(ctx, filter)
	if err != nil {
		log.Printf("error getting location: %v", err)
		return NewLocation(), err
	}

	var location LocationMongo
	if cursor.Next(ctx) {
		if err := cursor.Decode(&location); err != nil {
			return NewLocation(), err
		}
	}

	if cursor.Current == nil {
		log.Printf("location with element_type %s and element_id %s not found", elementType, elementId)
		return NewLocation(), ErrLocationNotFound
	}

	return &Location{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Altitude:  location.Altitude,
	}, nil

}
