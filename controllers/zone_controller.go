package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/models"
	"github.com/izacarias/lapi/responses"
	"github.com/izacarias/lapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var zoneCollection *mongo.Collection = configs.GetCollection(configs.DB, "zones")

// GetZone godoc
// @Router /queries/zones [get]
// @Tags location
// @Summary Query the information about one or more specific zones or a list of zones.
// @Description The GET method is used to query the information about one or more specific zones or a list of zones.
// @Id zonesGET
// @Produce json
// success 200 {array} ZoneResponse
func GetZone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var zone models.Zone
		zoneId := c.Param("id")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// create a filter document for the zone
		filter := bson.M{"id": zoneId}

		err := zoneCollection.FindOne(ctx, filter).Decode(&zone)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
				return
			}
			log.Printf("error getting zone %s: %v", zoneId, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching zone"})
			return
		}
		zr := responses.ZoneResponse{
			ZoneId:                            zone.Id,
			NumberOfAccessPoints:              0,
			NumberOfUnserviceableAccessPoints: 0,
			NumberOfUsers:                     0,
			ResourceURL:                       utils.ConstructZoneResourceUrl(c.Request, zone.Id),
		}
		// TODO: Process Zone Information here before returning
		c.JSON(http.StatusOK, zr)
	}
}

// TODO: Add zoneId as query parameter to filter zones (GS MEC 013 - 7.7.3.1-1)
func ListZones() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := zoneCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Printf("error getting zones: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching zones"})
			return
		}

		var zones []models.Zone
		if err = cursor.All(ctx, &zones); err != nil {
			log.Printf("error decoding zones: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching zones"})
			return
		}

		var zrs []responses.ZoneResponse
		for _, zone := range zones {
			zrs = append(zrs, responses.ZoneResponse{
				ZoneId:                            zone.Id,
				NumberOfAccessPoints:              0,
				NumberOfUnserviceableAccessPoints: 0,
				NumberOfUsers:                     0,
				ResourceURL:                       utils.ConstructZoneResourceUrl(c.Request, zone.Id),
			})
		}
		c.JSON(http.StatusOK, zrs)
	}
}
