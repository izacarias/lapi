package mock

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/domain/access_point"
	"github.com/izacarias/lapi/domain/zone"
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
			access_point.AccessPoint{
				AccessPointId:   "AP1",
				ConnectionType:  access_point.CT_WIFI,
				OperationStatus: access_point.OS_SERVICEABLE,
				NumberOfUsers:   10,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			access_point.AccessPoint{
				AccessPointId:   "AP2",
				ConnectionType:  access_point.CT_WIFI,
				OperationStatus: access_point.OS_SERVICEABLE,
				NumberOfUsers:   2,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			access_point.AccessPoint{
				AccessPointId:   "AP3",
				ConnectionType:  access_point.CT_5GNR,
				OperationStatus: access_point.OS_SERVICEABLE,
				NumberOfUsers:   15,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			access_point.AccessPoint{
				AccessPointId:   "AP4",
				ConnectionType:  access_point.CT_5GNR,
				OperationStatus: access_point.OS_SERVICEABLE,
				NumberOfUsers:   100,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			access_point.AccessPoint{
				AccessPointId:   "AP5",
				ConnectionType:  access_point.CT_5GNR,
				OperationStatus: access_point.OS_UNSERVICEABLE,
				NumberOfUsers:   0,
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			access_point.AccessPoint{
				AccessPointId:   "AP6",
				ConnectionType:  access_point.CT_WIFI,
				OperationStatus: access_point.OS_UNSERVICEABLE,
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
			zone.Zone{Id: "Zone1"},
			zone.Zone{Id: "Zone2"},
			zone.Zone{Id: "Zone3"},
			zone.Zone{Id: "Zone4"},
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
