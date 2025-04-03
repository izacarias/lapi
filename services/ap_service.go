package services

import (
	"log"

	"github.com/izacarias/lapi/domain"
)

// func GetZone(zoneId string) (*domain.Zone, error) {
// 	log.Printf("Getting zone with id: %v", zoneId)
// 	return domain.GetZone(zoneId)
// }

func ListApsInZone(zoneId string) ([]domain.AccessPoint, error) {
	log.Printf("getting all APs for zone %s", zoneId)
	zone, err := domain.GetZone(zoneId)
	if err != nil {
		return nil, err
	}
	return zone.GetAccessPoints(), nil
}

func GetApInZone(zoneId, apId string) (*domain.AccessPoint, error) {
	log.Printf("getting AP with id %s for zone %s", apId, zoneId)
	zone, err := domain.GetZone(zoneId)
	if err != nil {
		return nil, err
	}
	for _, ap := range zone.GetAccessPoints() {
		if ap.GetId() == apId {
			ap.SetLocation(getAPLocation(apId))
			return &ap, nil
		}
	}
	return nil, domain.ErrAccessPointNotFound
}

func getAPLocation(apId string) *domain.Location {
	log.Printf("getting location for AP %s", apId)
	location, err := domain.GetLocation(domain.TYPE_AP, apId)
	if err != nil {
		log.Printf("error getting location for AP %s: %v", apId, err)
	}
	return location
}
