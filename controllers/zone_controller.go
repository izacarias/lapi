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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var zoneCollection *mongo.Collection = configs.GetCollection(configs.DB, "zones")

func GetZone() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		zoneId := c.Param("id")
		var zone models.Zone
		defer cancel()

		objId, _ := zoneCollection.Find(ctx, bson.M{})

		err := zoneCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&zone)

		if err != nil {
			log.Printf("error getting zone %s: %v", zoneId, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching zone"})
			return
		}
		// TODO: Process Zone Information here before returning
		c.JSON(http.StatusOK, responses.ZoneResponse{
			ZoneId:                            zone.Id,
			NumberOfAccessPoints:              0,
			NumberOfUnserviceableAccessPoints: 0,
			NumberOfUsers:                     0,
			ResourceURL:                       "http://localhost:8080/api/v1/zone/" + zone.Id,
		})
	}
}
