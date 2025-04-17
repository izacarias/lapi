package services

import (
	"log"

	"github.com/izacarias/lapi/domain"
)

func GetZone(zoneId string) (*domain.Zone, error) {
	log.Printf("Getting zone with id: %v", zoneId)
	return domain.GetZone(zoneId)
}

func GetAllZones() ([]domain.Zone, error) {
	log.Printf("Getting all zones")
	return domain.GetAllZones()
}

func CountUsersInZone(zone *domain.Zone) int {
	zoneId := zone.GetId()
	log.Printf("Counting users in zone %s", zoneId)
	apsInZone, err := ListApsInZone(zoneId)
	if err != nil {
		log.Printf("error getting aps in zone %s: %v", zoneId, err)
		return 0
	}
	var useCount int = 0
	for _, ap := range apsInZone {
		useCount += ap.CountUsers()
	}
	return useCount
}
