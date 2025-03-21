package mock

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertMockData(client *mongo.Client) {
	// Insert mock data
	insertZoneData(client)
	insertAccessPointData(client)

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
			domain.ZoneMongo{
				ZoneId:       "zone1",
				AccessPoints: []string{"ap1", "ap2"},
			},
			domain.ZoneMongo{
				ZoneId:       "zone2",
				AccessPoints: []string{"ap3", "ap4"},
			},
			domain.ZoneMongo{
				ZoneId:       "zone3",
				AccessPoints: []string{"ap1", "ap2"},
			},
			domain.ZoneMongo{
				ZoneId:       "zone4",
				AccessPoints: []string{"ap3", "ap4"},
			},
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

func insertAccessPointData(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Get the Collection of Zones
	collection := client.Database("lapi").Collection("access_points")
	// Check if the collection is empty
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		mockData := []interface{}{
			domain.AccessPointMongo{
				ApId:            "ap1",
				ConnectionType:  string(domain.CT_WIFI),
				OperationStatus: string(domain.OS_SERVICEABLE),
				Users:           []string{"user1", "user2"},
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			domain.AccessPointMongo{
				ApId:            "ap2",
				ConnectionType:  string(domain.CT_WIFI),
				OperationStatus: string(domain.OS_SERVICEABLE),
				Users:           []string{"user3", "user4"},
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			domain.AccessPointMongo{
				ApId:            "ap3",
				ConnectionType:  string(domain.CT_5GNR),
				OperationStatus: string(domain.OS_SERVICEABLE),
				Users:           []string{"user5", "user6"},
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
			domain.AccessPointMongo{
				ApId:            "ap4",
				ConnectionType:  string(domain.CT_5GNR),
				OperationStatus: string(domain.OS_UNSERVICEABLE),
				Users:           []string{"user7", "user8"},
				Timezone:        "01-01-1970T00:00:00+01:00",
			},
		}

		_, err := collection.InsertMany(ctx, mockData)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Inserted mock data into the access_points collection")
	} else {
		log.Println("Access Points collection is not empty, skipping mock data insertion")
	}
}
