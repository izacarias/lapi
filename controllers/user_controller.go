package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

		c.String(http.StatusOK, "zoneId: %s, accessPointId: %s, userAddress: %s", qsZones, qsApId, qsUserAddress)
	}
}
