package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/responses"
	"github.com/izacarias/lapi/services"
)

func GetDistance() gin.HandlerFunc {
	return func(c *gin.Context) {

		qsUserAddress := c.QueryArray("address")
		qsLatituted := c.Query("latitude")
		qsLongitude := c.Query("longitude")

		// Possibilities
		// 1. Get distance between two users
		if len(qsUserAddress) == 2 {
			// TODO: sanitize user addresses
			userA, err := services.GetUserByAddress(qsUserAddress[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("User with address %s not found", qsUserAddress[0]),
					"status":  http.StatusBadRequest,
				})
				return
			}

			userB, err := services.GetUserByAddress(qsUserAddress[1])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("User with address %s not found", qsUserAddress[1]),
					"status":  http.StatusBadRequest,
				})
				return
			}

			calculatedDistance, err := services.CalculateDistance(userA, userB)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to calculate distance",
					"status":  http.StatusInternalServerError,
				})
				return
			}

			response := responses.TerminalDistance{
				Accuracy: calculatedDistance.GetAccuracy(),
				Distance: calculatedDistance.GetDistance(),
				Timestamp: &responses.TimeStamp{
					Seconds:     uint32(calculatedDistance.GetTimestamp()),
					NanoSeconds: 0,
				},
			}
			c.JSON(http.StatusOK, response)
			return
		}

		// 2. Get distance between a user and a coordinate (lat, lon)
		if len(qsUserAddress) == 1 && qsLatituted != "" && qsLongitude != "" {

		}

		// 3. Invalid query
		if len(qsUserAddress) > 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Query cannot have more than 2 'address' parameters",
				"status":  http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query must provide either 2 'address' parameters or 1 'address' parameter and 'latitude'/'longitude' parameters",
			"status":  http.StatusBadRequest,
		})

	}
}
