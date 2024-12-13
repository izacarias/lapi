package mongo

import (
	"context"
	"time"

	"github.com/izacarias/lapi/domain/access_point"
	"github.com/izacarias/lapi/domain/zone"
	"github.com/izacarias/lapi/domain/zone_association"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Database
	// store the zones and associated access points
	zone *mongo.Collection
}

type MongoAssociation struct {
	id string
	accessPoints []string
}

func (mr *MongoRepository) Get(zoneId string) (zone_association.ZoneAssociation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.zone.FindOne(ctx, bson.M{"id": zoneId})

	var z zone.Zone
	if err := result.Decode(&z); err != nil {
		return zone_association.ZoneAssociation{}, err
	}
	var aps []access_point.AccessPoint
	for _, apId := range  {
		// get the access point
	}




	return zone.Zone{}, nil
}
