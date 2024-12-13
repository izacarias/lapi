package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
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

// ListZones godoc
// @Tags location
// @Summary Query the information about one or more specific zones or a list of zones.
// @Description The GET method is used to query the information about one or more specific zones or a list of zones.
// @Id zonesGET
// @Produce json
// @Param zoneId query []string false "Zone ID"
// @Success 200 {object} responses.ZoneList
// @Failure 500
// @Router /queries/zones [get]
func ListZones() gin.HandlerFunc {
	return func(c *gin.Context) {
		var shouldFilter bool
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// get the zone ids from the query parameters
		zoneIds := c.QueryArray("zoneId")
		if zoneIds != nil {
			zoneIds = strings.Split(zoneIds[0], ",")
			shouldFilter = len(zoneIds) > 0
		}

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

		var zrs responses.ZoneList
		zrs.Zone = make([]responses.ZoneInfo, 0)
		if shouldFilter {
			for _, zone := range zones {
				for _, zoneId := range zoneIds {
					if zone.Id == zoneId {
						zrs.Zone = append(zrs.Zone, responses.ZoneInfo{
							ZoneId:                            zone.Id,
							NumberOfAccessPoints:              int32(zone.CountAccessPoints()),
							NumberOfUnserviceableAccessPoints: 0,
							NumberOfUsers:                     0,
							ResourceURL:                       utils.ConstructZoneResourceUrl(c.Request, zone.Id),
						})
					}
				}
			}
		} else {
			for _, zone := range zones {
				zrs.Zone = append(zrs.Zone, responses.ZoneInfo{
					ZoneId:                            zone.Id,
					NumberOfAccessPoints:              int32(zone.CountAccessPoints()),
					NumberOfUnserviceableAccessPoints: 0,
					NumberOfUsers:                     0,
					ResourceURL:                       utils.ConstructZoneResourceUrl(c.Request, zone.Id),
				})
			}
		}
		c.JSON(http.StatusOK, zrs)
	}
}

// GetZone godoc
// @Tags location
// @Summary Query information about a specific zone
// @Description The GET method is used to query the information about a specific zone.
// @Id zoneGetById
// @Produce json
// @Param	zoneId	path	string	true	"Zone ID"
// @Success	200 {object}	responses.ZoneInfo
// @Failure	400	"Bad Request"
// @Failure	404	"Not found"

// @Failure	500	"Internal Server Error"
// @Router /queries/zones/{zoneId} [get]
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
		zr := responses.ZoneInfo{
			ZoneId:                            zone.Id,
			NumberOfAccessPoints:              int32(zone.CountAccessPoints()),
			NumberOfUnserviceableAccessPoints: 0,
			NumberOfUsers:                     0,
			ResourceURL:                       utils.ConstructZoneResourceUrl(c.Request, zone.Id),
		}
		// TODO: Process Zone Information here before returning
		c.JSON(http.StatusOK, zr)
	}
}
