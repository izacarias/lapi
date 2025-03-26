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

func ListAccessPoints() gin.HandlerFunc {
	return func(c *gin.Context) {
		zoneId := c.Param("id")

		if zoneId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "zone id is required"})
			return
		}
		aps, err := services.GetApsInZone(zoneId)
		if err != nil {
			log.Printf("error getting zone %s: %v", zoneId, err)

			if err == domain.ErrZoneNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching zone"})
			return
		}
		apsResponse := make([]responses.AccessPointInfo, 0)
		for _, ap := range aps {
			apsResponse = append(apsResponse, responses.AccessPointInfo{
				AccessPointId:   ap.GetId(),
				OperationStatus: responses.OperationStatus(ap.GetOperationStatus()),
				NumberOfUsers:   int32(ap.CountUsers()),
				ResourceURL:     utils.GetAccessPointResourceUrl(c.Request, zoneId, ap.GetId()),
			})
		}

		response := responses.AccessPointList{
			AccessPoint: responses.AccessPointInfoList{
				ZoneId:      zoneId,
				AccessPoint: apsResponse,
				ResourceURL: utils.GetAccessPointListResourceUrl(c.Request, zoneId),
			},
		}
		c.JSON(http.StatusOK, response)
	}
}
