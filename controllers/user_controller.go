package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/domain"
	"github.com/izacarias/lapi/responses"
	"github.com/izacarias/lapi/services"
	"github.com/izacarias/lapi/utils"
)

func ListUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		qsZones := c.QueryArray("zoneId")
		qsApId := c.QueryArray("accessPointId")
		qsUserAddress := c.QueryArray("address")

		// qsUserAddress is not empty
		if len(qsUserAddress) > 0 {
			userList := getUsersByAddresses(qsUserAddress)

			// Apply filters based on query parameters
			userList = applyFilters(userList, qsZones, qsApId)

			// Convert to response format and return
			response := formatUserResponse(c, userList, qsUserAddress[0])
			c.JSON(http.StatusOK, response)
			return
		} else {
			userList, err := services.GetAllUsers()
			if err != nil {
				log.Printf("Error retrieving all users: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
				return
			}
			// TODO: filter by zoneId and/or accessPointId
			userList = applyFilters(userList, qsZones, qsApId)
			// Convert to response format and return
			response := formatUserResponse(c, userList, "")
			c.JSON(http.StatusOK, response)
			return
		}
	}
}

// getUsersByAddresses retrieves users by their addresses
func getUsersByAddresses(addresses []string) []domain.User {
	userList := make([]domain.User, 0)
	for _, userAddress := range addresses {
		user, err := services.GetUserByAddress(userAddress)
		if err != nil {
			// Note: This should be handled better - possibly returning an error
			// Currently preserving the original behavior
			log.Printf("Error retrieving user by address %s: %v", userAddress, err)
			continue
		}
		userList = append(userList, *user)
	}
	return userList
}

// applyFilters applies zone and access point filters to the user list
func applyFilters(userList []domain.User, qsZones []string, qsApId []string) []domain.User {
	// filter by Zone only
	if len(qsZones) > 0 && len(qsApId) == 0 {
		return filterByZone(userList, qsZones)
	}

	// filter by Access Point only
	if len(qsZones) == 0 && len(qsApId) > 0 {
		return filterByAccessPoint(userList, qsApId)
	}

	// filter by both Zone and Access Point
	if len(qsZones) > 0 && len(qsApId) > 0 {
		return filterByZoneAndAccessPoint(userList, qsZones, qsApId)
	}

	// No filters applied
	return userList
}

// filterByZone filters users by zone IDs
func filterByZone(userList []domain.User, zoneIds []string) []domain.User {
	userListFiltered := make([]domain.User, 0)
	for _, zoneId := range zoneIds {
		for _, user := range userList {
			if user.ZoneId == zoneId {
				userListFiltered = append(userListFiltered, user)
			}
		}
	}
	return userListFiltered
}

// filterByAccessPoint filters users by access point IDs
func filterByAccessPoint(userList []domain.User, accessPointIds []string) []domain.User {
	userListFiltered := make([]domain.User, 0)
	for _, accessPointId := range accessPointIds {
		for _, user := range userList {
			if user.AccessPoint == accessPointId {
				userListFiltered = append(userListFiltered, user)
			}
		}
	}
	return userListFiltered
}

// filterByZoneAndAccessPoint filters users by both zone and access point IDs
func filterByZoneAndAccessPoint(userList []domain.User, zoneIds []string, accessPointIds []string) []domain.User {
	userListFiltered := make([]domain.User, 0)
	for _, zoneId := range zoneIds {
		for _, accessPointId := range accessPointIds {
			for _, user := range userList {
				if user.ZoneId == zoneId && user.AccessPoint == accessPointId {
					userListFiltered = append(userListFiltered, user)
				}
			}
		}
	}
	return userListFiltered
}

// formatUserResponse converts domain users to response format
func formatUserResponse(c *gin.Context, userList []domain.User, queryAddress string) responses.UserInfoList {
	usersResponse := make([]responses.UserInfo, 0)
	for _, user := range userList {
		usersResponse = append(usersResponse, responses.UserInfo{
			Address:       user.Address,
			AccessPointId: &user.AccessPoint,
			ZoneId:        user.ZoneId,
			ResourceURL:   utils.GetUserResourceUrl(c.Request, user.Address),
			//TODO: Update with the timestamp from the last location update
			Timestamp: responses.TimeStamp{Seconds: uint32(100), NanoSeconds: uint32(200)},
			//TODO: Update with the location info
			LocationInfo: &responses.LocationInfo{
				Latitude:  []float32{user.Location.Latitude},
				Longitude: []float32{user.Location.Longitude},
				Altitude:  user.Location.Altitude,
				Shape:     responses.LocationInfoShapeN2,
			},
		})
	}

	return responses.UserInfoList{
		ResourceURL: utils.GetUserResourceUrl(c.Request, queryAddress),
		User:        usersResponse,
	}
}
