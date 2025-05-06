package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/responses"
	"github.com/izacarias/lapi/services"
)

// GetDistance godoc
// @Summary Query information about distance from a user to a location or between two users
// @Description The GET method is used to query information about distance from a user to a location or between two users.
// @Id distanceGET
// @Tags location
// @Produce json
// @Param address query []string true "User address" // 2 addresses for distance between users
// @Param latitude query string false "Latitude" // 1 address and lat/lon for distance to a location
// @Param longitude query string false "Longitude" // 1 address and lat/lon for distance to a location
// @Success 200 {object} responses.TerminalDistance
// @Failure 400 {object} responses.ProblemDetails "Bad Request"
// @Failure 404 {object} responses.ProblemDetails "User Not Found"
// @Failure 500 {object} responses.ProblemDetails "Internal Server Error"
// @Router /queries/distance [get]
func GetDistance() gin.HandlerFunc {
	return func(c *gin.Context) {

		qsUserAddress := c.QueryArray("address")
		qsLatitude := c.Query("latitude")
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
					Seconds:     uint32(calculatedDistance.GetTimestamp().Unix()),
					NanoSeconds: 0,
				},
			}
			c.JSON(http.StatusOK, response)
			return
		}

		// 2. Get distance between a user and a coordinate (lat, lon)
		if len(qsUserAddress) == 1 && qsLatitude != "" && qsLongitude != "" {
			userA, err := services.GetUserByAddress(qsUserAddress[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": fmt.Sprintf("User with address %s not found", qsUserAddress[0]),
					"status":  http.StatusBadRequest,
				})
				return
			}
			calculatedDistance, err := services.CalculateDistanceLatLong(userA, qsLatitude, qsLongitude)

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
					Seconds:     uint32(calculatedDistance.GetTimestamp().Unix()),
					NanoSeconds: 0,
				},
			}
			c.JSON(http.StatusOK, response)
			return

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
