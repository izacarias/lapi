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
	ZoneId string `json:"zone_id" bson:"zone_id"`
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
	aps, err := getAccessPointsByZoneId(zoneId)
	if err != nil {
		log.Printf("error getting access points: %v", err)
		return NewZone(), err
	}
	for _, ap := range aps {
		zone.AddAccessPoint(ap)
	}
	return zone, nil
}

func GetAllZones() ([]Zone, error) {
	zones := make([]Zone, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := zoneCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("error getting zones: %v", err)
		return nil, err
	}

	var zonesDb []ZoneMongo
	if err = cursor.All(ctx, &zonesDb); err != nil {
		log.Printf("error decoding zones: %v", err)
		return nil, err
	}
	// iterate over all elements in zoneDb and create a Zone object
	for _, z := range zonesDb {
		zone := NewZone()
		zone.SetId(z.ZoneId)
		// add access points to the zone
		aps, err := getAccessPointsByZoneId(z.ZoneId)
		if err != nil {
			log.Printf("error getting access points: %v", err)
		}
		for _, ap := range aps {
			zone.AddAccessPoint(ap)
		}
		zones = append(zones, *zone)
	}
	return zones, nil
}
