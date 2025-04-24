package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/domain"
	"github.com/izacarias/lapi/services"
)

type mininetUserLocationUpdate struct {
	UserID    string  `json:"userid"`
	CoordX    float64 `json:"coordx"`
	CoordY    float64 `json:"coordy"`
	CoordZ    float64 `json:"coordz"`
	APName    string  `json:"apname"`
	Timestamp int     `json:"timestamp"`
}

type mininetApLocationUpdate struct {
	ApId      string  `json:"apid"`
	CoordX    float64 `json:"coordx"`
	CoordY    float64 `json:"coordy"`
	CoordZ    float64 `json:"coordz"`
	Timestamp int     `json:"timestamp"`
}

func UpdateUserLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the JSON request body
		var locationMn mininetUserLocationUpdate
		if err := c.ShouldBindJSON(&locationMn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
			log.Printf("Error parsing JSON: %v", err)
			return
		}
		log.Printf("Received location update: %+v", locationMn)
		location := domain.NewLocation()
		timestamp := time.Unix(int64(locationMn.Timestamp), 0)

		location.SetLatitude(float32(locationMn.CoordX))
		location.SetLongitude(float32(locationMn.CoordY))
		location.SetAltitude(float32(locationMn.CoordZ))
		location.SetTimestamp(timestamp)
		services.UpdateUserLocation(locationMn.UserID, locationMn.APName, location)

		c.JSON(http.StatusOK, gin.H{"status": "Location updated successfully", "location": locationMn})
	}
}

func UpdateAPLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the JSON request body
		var mnAP mininetApLocationUpdate
		if err := c.ShouldBindJSON(&mnAP); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
			log.Printf("Error parsing JSON: %v", err)
			return
		}
		log.Printf("Received location update for AP: %+v", mnAP)
		location := domain.NewLocation()
		timestamp := time.Unix(int64(mnAP.Timestamp), 0)

		location.SetLatitude(float32(mnAP.CoordX))
		location.SetLongitude(float32(mnAP.CoordY))
		location.SetAltitude(float32(mnAP.CoordZ))
		location.SetTimestamp(timestamp)
		services.UpdateAPLocation(mnAP.ApId, location)

		c.JSON(http.StatusOK, gin.H{"status": "Location updated successfully", "location": mnAP})
	}
}
