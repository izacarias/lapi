package mock

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertMockData(client *mongo.Client) {
	// Inserting mock data
	collection := client.Database("lapi").Collection("zones")

	// Check if the collection is empty
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		mockData := []interface{}{
			models.Zone{Id: "Zone1"},
			models.Zone{Id: "Zone2"},
			models.Zone{Id: "Zone3"},
			models.Zone{Id: "Zone4"},
		}

		_, err := collection.InsertMany(ctx, mockData)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Inserted mock data into the zones collection")
	} else {
		log.Println("Zones collection is not empty, skipping mock data insertion")
	}
}
