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

		// get the zone ids from the query parameters
		// zoneIds := c.QueryArray("zoneId")
		// if zoneIds != nil {
		// 	zoneIds = strings.Split(zoneIds[0], ",")
		// }

		zones, err := services.GetAllZones()

		if err != nil {
			log.Printf("error getting zones: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching zones"})
			return
		}

		var zoneInfoList = make([]responses.ZoneInfo, 0)

		for _, zone := range zones {
			zoneInfoList = append(zoneInfoList, responses.ZoneInfo{
				ZoneId:                            zone.GetId(),
				NumberOfAccessPoints:              int32(zone.CountAccessPoints()),
				NumberOfUnserviceableAccessPoints: int32(zone.CountSericeableAccessPoints()),
				NumberOfUsers:                     int32(zone.CountUsersInZone()),
				ResourceURL:                       utils.GetZoneResourceUrl(c.Request, zone.GetId()),
			})
		}
		response := responses.ZoneList{
			ZoneList: responses.ZoneInfoList{
				ResourceURL: utils.GetZoneListResourceUrl(c.Request),
				Zone:        zoneInfoList,
			},
		}
		c.JSON(http.StatusOK, response)
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
		zoneId := c.Param("id")

		// issue error 400 of zoneId is empty
		if zoneId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "zoneId is required"})
			return
		}

		zone, err := services.GetZone(zoneId)

		if err != nil {
			log.Printf("error getting zone %s: %v", zoneId, err)

			if err == domain.ErrZoneNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching zone"})
			return
		}

		zr := responses.ZoneInfo{
			ZoneId:                            zone.GetId(),
			NumberOfAccessPoints:              int32(zone.CountAccessPoints()),
			NumberOfUnserviceableAccessPoints: int32(zone.CountSericeableAccessPoints()),
			NumberOfUsers:                     int32(zone.CountUsersInZone()),
			ResourceURL:                       utils.GetZoneResourceUrl(c.Request, zone.GetId()),
		}
		// TODO: Process Zone Information here before returning
		c.JSON(http.StatusOK, zr)
	}
}
