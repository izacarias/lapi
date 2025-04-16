package mock

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName          = "lapi"
	contextTimeout  = 10 * time.Second
	zonesCollection = "zones"
	apsCollection   = "access_points"
	locsCollection  = "locations"
	usersCollection = "users"
)

func InsertMockData(client *mongo.Client) {
	// Clear the database if configured
	if configs.GetConfigClearDatabase() {
		clearDatabase(client)
	}

	// Skip if mock data insertion is disabled
	if !configs.GetConfigInsertMockData() {
		log.Println("Insert mock data is disabled, skipping")
		return
	}

	// Insert mock data for each collection
	insertZoneData(client)
	insertAccessPointData(client)
	insertLocationData(client)
	insertUserData(client)
}

// clearDatabase drops all collections in the database
func clearDatabase(client *mongo.Client) {
	ctx, cancel := getContext()
	defer cancel()

	// Get all collections in the database
	collections, err := client.Database(dbName).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Drop each collection
	for _, collection := range collections {
		if err := client.Database(dbName).Collection(collection).Drop(ctx); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Database cleared successfully")
}

// insertCollectionData is a generic function to insert data into a collection if it's empty
func insertCollectionData(client *mongo.Client, collectionName string, mockData []interface{}, description string) {
	ctx, cancel := getContext()
	defer cancel()

	collection := client.Database(dbName).Collection(collectionName)

	// Check if collection is empty
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		_, err := collection.InsertMany(ctx, mockData)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Inserted mock data into the %s collection", description)
	} else {
		log.Printf("%s collection is not empty, skipping mock data insertion", description)
	}
}

// getContext creates a context with timeout
func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), contextTimeout)
}

// insertZoneData adds mock zones to the database
func insertZoneData(client *mongo.Client) {
	mockData := []interface{}{
		domain.ZoneMongo{ZoneId: "zone1"},
		domain.ZoneMongo{ZoneId: "zone2"},
		domain.ZoneMongo{ZoneId: "zone3"},
		domain.ZoneMongo{ZoneId: "zone4"},
	}

	insertCollectionData(client, zonesCollection, mockData, "zones")
}

// insertAccessPointData adds mock access points to the database
func insertAccessPointData(client *mongo.Client) {
	mockData := []interface{}{
		domain.AccessPointMongo{
			ApId:            "ap1",
			ConnectionType:  string(domain.CT_WIFI),
			OperationStatus: string(domain.OS_SERVICEABLE),
			Timezone:        "01-01-1970T00:00:00+01:00",
			ZoneId:          "zone1",
		},
		domain.AccessPointMongo{
			ApId:            "ap2",
			ConnectionType:  string(domain.CT_WIFI),
			OperationStatus: string(domain.OS_SERVICEABLE),
			Timezone:        "01-01-1970T00:00:00+01:00",
			ZoneId:          "zone2",
		},
		domain.AccessPointMongo{
			ApId:            "ap3",
			ConnectionType:  string(domain.CT_5GNR),
			OperationStatus: string(domain.OS_SERVICEABLE),
			Timezone:        "01-01-1970T00:00:00+01:00",
			ZoneId:          "zone3",
		},
		domain.AccessPointMongo{
			ApId:            "ap4",
			ConnectionType:  string(domain.CT_5GNR),
			OperationStatus: string(domain.OS_UNSERVICEABLE),
			Timezone:        "01-01-1970T00:00:00+01:00",
			ZoneId:          "zone4",
		},
	}

	insertCollectionData(client, apsCollection, mockData, "access_points")
}

// insertLocationData adds mock location data to the database
func insertLocationData(client *mongo.Client) {
	currentTime := time.Now().Unix()
	oldTime := currentTime - 1000000
	mockData := []interface{}{
		domain.LocationMongo{
			ElementType: domain.TYPE_AP,
			ElementId:   "ap1",
			Latitude:    1.0,
			Longitude:   1.0,
			Altitude:    1.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_AP,
			ElementId:   "ap2",
			Latitude:    2.0,
			Longitude:   2.0,
			Altitude:    2.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.1",
			Latitude:    3.0,
			Longitude:   3.0,
			Altitude:    3.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.2",
			Latitude:    4.0,
			Longitude:   4.0,
			Altitude:    4.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.3",
			Latitude:    5.0,
			Longitude:   5.0,
			Altitude:    5.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.4",
			Latitude:    6.0,
			Longitude:   6.0,
			Altitude:    6.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.5",
			Latitude:    5.0,
			Longitude:   5.0,
			Altitude:    6.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.5",
			Latitude:    10.0,
			Longitude:   10.0,
			Altitude:    10.0,
			Timestamp:   oldTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.6",
			Latitude:    6.0,
			Longitude:   6.0,
			Altitude:    6.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.7",
			Latitude:    7.0,
			Longitude:   7.0,
			Altitude:    7.0,
			Timestamp:   currentTime,
		},
		domain.LocationMongo{
			ElementType: domain.TYPE_USER,
			ElementId:   "192.168.1.8",
			Latitude:    6.0,
			Longitude:   6.0,
			Altitude:    6.0,
			Timestamp:   currentTime,
		},
	}

	insertCollectionData(client, locsCollection, mockData, "locations")
}

// insertUserData adds mock user data to the database
func insertUserData(client *mongo.Client) {
	mockData := []interface{}{
		domain.UserMongo{Address: "192.168.1.1", AccessPoint: "ap1"},
		domain.UserMongo{Address: "192.168.1.2", AccessPoint: "ap2"},
		domain.UserMongo{Address: "192.168.1.3", AccessPoint: "ap3"},
		domain.UserMongo{Address: "192.168.1.4", AccessPoint: "ap4"},
		domain.UserMongo{Address: "192.168.1.5", AccessPoint: "ap1"},
		domain.UserMongo{Address: "192.168.1.6", AccessPoint: "ap2"},
		domain.UserMongo{Address: "192.168.1.7", AccessPoint: "ap1"},
		domain.UserMongo{Address: "192.168.1.8", AccessPoint: "ap1"},
	}

	insertCollectionData(client, usersCollection, mockData, "users")
}
