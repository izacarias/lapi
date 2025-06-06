package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/domain"
	"github.com/izacarias/lapi/services"
)

type mininetUser struct {
	UserID string `json:"userid"`
	APName string `json:"apname"`
}

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

// RegisterUser godoc
// @Summary Register users from Mininet into the local database
// @Description The POST method is used to register a new users within the location service
// @Id registerMnUserPOST
// @Tags mininet
// @Accept json
// @Produce json
// @Param user body mininetUser true "User data"
// @Success 200 {object} string "User registered successfully"
// @Failure 400 {object} responses.ProblemDetails "Bad Request"
// @Failure 500 {object} responses.ProblemDetails "Internal Server Error"
// @Router /mininet/user [post]
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user mininetUser
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format: " + err.Error()})
			log.Printf("Error parsing JSON: %v", err)
			return
		}
		log.Printf("Received user registration data: %+v", user)
		newUser := domain.NewUser()
		newUser.SetAddress(user.UserID)
		newUser.SetAccessPoint(user.APName)
		services.InsertUser(newUser)
		c.JSON(http.StatusOK, gin.H{"status": "User registered successfully", "user": newUser})
	}
}

// UpdateUserLocation godoc
// @Summary Update the location of a user with data from Mininet
// @Description The POST method is used to update the location of a user
// @Id updateMnUserLocationPOST
// @Tags mininet
// @Accept json
// @Produce json
// @Param location body mininetUserLocationUpdate true "User Location Update"
// @Success 200 {object} string "Location updated successfully"
// @Failure 400 {object} responses.ProblemDetails "Bad Request"
// @Failure 500 {object} responses.ProblemDetails "Internal Server Error"
// @Router /mininet/location [post]
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

func RegisterAccessPoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		// This function is currently not implemented
		c.JSON(http.StatusNotImplemented, gin.H{"error": "This endpoint is not implemented yet"})
	}
}

// UpdateAPLocation godoc
// @Summary Update the location of an access point with data from Mininet
// @Description The POST method is used to update the location of an access point
// @Id updateMnAPLocationPOST
// @Tags mininet
// @Accept json
// @Produce json
// @Param location body mininetApLocationUpdate true "Access Point Location Update"
// @Success 200 {object} string "Location updated successfully"
// @Failure 400 {object} responses.ProblemDetails "Bad Request"
// @Failure 500 {object} responses.ProblemDetails "Internal Server Error"
// @Router /mininet/aplocation [post]
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
