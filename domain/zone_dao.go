package domain

import (
	"context"
	"log"
	"time"

	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var zoneCollection *mongo.Collection = configs.GetCollection(configs.DB, "zones")

type ZoneMongo struct {
	ZoneId       string   `json:"zone_id" bson:"zone_id"`
	AccessPoints []string `json:"access_points" bson:"access_points"`
}

func GetZone(zoneId string) (*Zone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter the zones by the zoneId
	filter := bson.M{"zone_id": zoneId}
	cursor, err := zoneCollection.Find(ctx, filter)
	if err != nil {
		log.Printf("error getting zones: %v", err)
		return NewZone(), err
	}

	var z ZoneMongo
	if cursor.Next(ctx) {
		if err := cursor.Decode(&z); err != nil {
			return NewZone(), err
		}
	}

	if cursor.Current == nil {
		log.Printf("zone with id %v not found", zoneId)
		return NewZone(), ErrZoneNotFound
	}

	zone := NewZone()
	zone.SetId(z.ZoneId)

	return zone, nil
}
