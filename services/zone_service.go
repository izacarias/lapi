package services

import (
	"log"

	"github.com/izacarias/lapi/domain"
)

func GetZone(zoneId string) (*domain.Zone, error) {
	log.Printf("Getting zone with id: %v", zoneId)
	return domain.GetZone(zoneId)
}
