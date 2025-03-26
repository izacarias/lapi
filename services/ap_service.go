package services

import (
	"log"

	"github.com/izacarias/lapi/domain"
)

// func GetZone(zoneId string) (*domain.Zone, error) {
// 	log.Printf("Getting zone with id: %v", zoneId)
// 	return domain.GetZone(zoneId)
// }

func GetApsInZone(zoneId string) ([]domain.AccessPoint, error) {
	log.Printf("Getting all APs for zone %s", zoneId)
	zone, err := domain.GetZone(zoneId)
	if err != nil {
		return nil, err
	}
	return zone.GetAccessPoints(), nil
}
