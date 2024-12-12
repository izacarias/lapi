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
	// Insert mock data
	insertZoneData(client)
	insertAccessPointData(client)

}

func insertAccessPointData(client *mongo.Client) {
	collectionName := "access_points"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Get the Collection of Zones
	collection := client.Database("lapi").Collection(collectionName)
	// Check if the collection is empty
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		mockData := []interface{}{
			models.AccessPoint{
				AccessPointId:   "AP1",
				ConnectionType:  models.CT_WIFI,
				OperationStatus: models.OS_SERVICEABLE,
				NumberOfUsers:   10,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			models.AccessPoint{
				AccessPointId:   "AP2",
				ConnectionType:  models.CT_WIFI,
				OperationStatus: models.OS_SERVICEABLE,
				NumberOfUsers:   2,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			models.AccessPoint{
				AccessPointId:   "AP3",
				ConnectionType:  models.CT_5GNR,
				OperationStatus: models.OS_SERVICEABLE,
				NumberOfUsers:   15,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			models.AccessPoint{
				AccessPointId:   "AP4",
				ConnectionType:  models.CT_5GNR,
				OperationStatus: models.OS_SERVICEABLE,
				NumberOfUsers:   100,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			models.AccessPoint{
				AccessPointId:   "AP5",
				ConnectionType:  models.CT_5GNR,
				OperationStatus: models.OS_UNSERVICEABLE,
				NumberOfUsers:   0,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			models.AccessPoint{
				AccessPointId:   "AP6",
				ConnectionType:  models.CT_WIFI,
				OperationStatus: models.OS_UNSERVICEABLE,
				NumberOfUsers:   0,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
		}

		_, err := collection.InsertMany(ctx, mockData)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("inserted mock data into the %s collection", collectionName)
	} else {
		log.Printf("%s collection is not empty, skipping mock data insertion", collectionName)
	}
}

func insertZoneData(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Get the Collection of Zones
	collection := client.Database("lapi").Collection("zones")
	// Check if the collection is empty
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		mockData := []interface{}{
			models.Zone{Id: "Zone1", AccessPointsList: []string{"AP1", "AP2"}},
			models.Zone{Id: "Zone2", AccessPointsList: []string{"AP3", "AP4"}},
			models.Zone{Id: "Zone3", AccessPointsList: []string{"AP5", "AP6"}},
			models.Zone{Id: "Zone4", AccessPointsList: []string{"AP1", "AP2", "AP5", "AP6"}},
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
