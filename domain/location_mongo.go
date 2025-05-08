package domain

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLocationRepository struct {
	// MongoDB client
	client             *mongo.Client
	locationCollection *mongo.Collection
}

type mongoLocation struct {
	ElementType string  `bson:"element_type"`
	ElementId   string  `bson:"element_id"`
	Latitude    float32 `bson:"latitude"`
	Longitude   float32 `bson:"longitude"`
	Altitude    float32 `bson:"altitude"`
	Timestamp   int64   `bson:"timestamp"`
}

func (ml *mongoLocation) ToDomain() *Location {
	l := Location{}
	l.Latitude = ml.Latitude
	l.Longitude = ml.Longitude
	l.Altitude = ml.Altitude
	l.Timestamp = time.Unix(ml.Timestamp, 0)
	return &l
}

func NewLocationMongo() (*MongoLocationRepository, error) {
	client := configs.DB
	collection := configs.GetCollection(client, "locations")

	return &MongoLocationRepository{
		client:             client,
		locationCollection: collection,
	}, nil
}

func (mr *MongoLocationRepository) Add(elementType string, elementId string, location *Location) error {
	return nil
}

func (mr *MongoLocationRepository) GetLast(elementType string, elementId string) (*Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Set up options to sort by timestamp in descending order (most recent first)
	opts := options.Find().SetSort(bson.M{"timestamp": -1})
	filter := bson.M{"element_type": elementType, "element_id": elementId}
	cursor, err := mr.locationCollection.Find(ctx, filter, opts)
	if err != nil {
		log.Printf("error getting location: %v", err)
		return nil, err
	}

	var ml mongoLocation
	if cursor.Next(ctx) {
		if err := cursor.Decode(&ml); err != nil {
			return nil, err
		}
	}

	if cursor.Current == nil {
		log.Printf("location with element_type %s and element_id %s not found", elementType, elementId)
		return nil, ErrLocationNotFound
	}

	return ml.ToDomain(), nil
}
