package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/domain"
	"github.com/izacarias/lapi/services"
)

/* Expected response
{
  "userList": {
    "resourceURL": "https://try-mec.etsi.org/sbxoyur055/mep1/location/v3/queries/users",
    "user": [
      {
        "address": "10.100.0.1",
        "accessPointId": "4g-macro-cell-3",
        "zoneId": "zone01",
        "resourceURL": "https://try-mec.etsi.org/sbxoyur055/mep1/location/v3/queries/users?address=10.100.0.1",
        "timestamp": {
          "nanoSeconds": 0,
          "seconds": 1743697065
        },
        "locationInfo": {
          "latitude": [
            43.735428
          ],
          "longitude": [
            7.417364
          ],
          "shape": 2
        },
        "civicInfo": {
          "country": "MC"
        },
        "relativeLocationInfo": {
          "X": 447.71487,
          "Y": -315.28012,
          "mapInfo": {
            "mapId": "324561243",
            "origin": {
              "latitude": 43.7314,
              "longitude": 7.4202
            }
          }
        }
      }
    ]
  }
}

*/

func ListUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		qsZones := c.QueryArray("zoneId")
		qsApId := c.QueryArray("accessPointId")
		qsUserAddress := c.QueryArray("userAddress")
		// possible workflows

		// qsUserAddress is not empty
		if len(qsUserAddress) > 0 {
			// filter by user address
			userList := make([]domain.User, 0)
			for _, userAddress := range qsUserAddress {
				user, err := services.GetUserByAddress(userAddress)
				if err != nil {
					c.String(http.StatusInternalServerError, "error getting user by address: %s", err)
					return
				}
				if len(qsZones) > 0 && len(qsApId) == 0 {
					// if qsZones is not empty, only add users with matching zoneId

				}
				if len(qsZones) == 0 && len(qsApId) > 0 {
					// if qsApId is not empty, only add users with matching accessPointId

				}
				if len(qsZones) > 0 && len(qsApId) > 0 {
					// if both qsZones and qsApId are not empty, only add users with matching zoneId and accessPointId

				}
				// if both qsZones and qsApId are empty, add all users
				if len(qsZones) == 0 && len(qsApId) == 0 {
					// if both qsZones and qsApId are not empty, only add users with matching zoneId and accessPointId
					userList = append(userList, *user)
				}

			}

		} else {
			// TODO: filter by zoneId and/or accessPointId
			if len(qsZones) > 0 && len(qsApId) == 0 {
				// get all users in the zone
			}
			if len(qsZones) == 0 && len(qsApId) > 0 {
				// get all users in the access point
			}
			if len(qsZones) > 0 && len(qsApId) > 0 {
				// get all users in the zone and filter for users in the provided access point list
				//TODO: get all users in the zone and filter for users in the provided access point list
			}
		}

		c.String(http.StatusOK, "zoneId: %s, accessPointId: %s, userAddress: %s", qsZones, qsApId, qsUserAddress)
	}
}
