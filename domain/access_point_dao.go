package domain

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accessPointCollection *mongo.Collection = configs.GetCollection(configs.DB, "access_points")

type AccessPointMongo struct {
	ApId            string   `bson:"access_point_id"`
	ConnectionType  string   `bson:"connection_type"`
	OperationStatus string   `bson:"operation_status"`
	Timezone        string   `bson:"timezone"`
	Users           []string `bson:"users"`
}

func GetAccessPoint(apId string) (*AccessPoint, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter ap by the apId
	filter := bson.M{"access_point_id": apId}
	cursor, err := accessPointCollection.Find(ctx, filter)
	if err != nil {
		log.Printf("error getting access point: %v", err)
		return NewAccessPoint(), err
	}

	var ap AccessPointMongo
	if cursor.Next(ctx) {
		if err := cursor.Decode(&ap); err != nil {
			return NewAccessPoint(), err
		}
	}

	if cursor.Current == nil {
		log.Printf("access point with id %v not found", apId)
		return NewAccessPoint(), ErrAccessPointNotFound
	}

	accessPoint := NewAccessPoint()
	accessPoint.SetId(ap.ApId)
	accessPoint.SetConnectionType(ConnectionType(ap.ConnectionType))
	accessPoint.SetOperationStatus(OperationStatus(ap.OperationStatus))
	accessPoint.SetTimeZone(ap.Timezone)
	// add users to the access point
	for _, userId := range ap.Users {
		user := NewUser(userId)
		accessPoint.AddUser(user)
	}
	return accessPoint, nil
}
