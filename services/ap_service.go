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

	zone, err := getZoneById(zoneId)
	if err != nil {
		return nil, err
	}

	return enrichAccessPointsWithLocation(zone.GetAccessPoints()), nil
}

func GetApInZone(zoneId, apId string) (*domain.AccessPoint, error) {
	log.Printf("getting AP with id %s for zone %s", apId, zoneId)

	zone, err := getZoneById(zoneId)
	if err != nil {
		return nil, err
	}

	for _, ap := range zone.GetAccessPoints() {
		if ap.GetId() == apId {
			enrichedAp := enrichAccessPointWithLocation(ap)
			return &enrichedAp, nil
		}
	}

	return nil, domain.ErrAccessPointNotFound
}

// getZoneById retrieves a zone by its ID
func getZoneById(zoneId string) (*domain.Zone, error) {
	zone, err := domain.GetZone(zoneId)
	if err != nil {
		log.Printf("error getting zone %s: %v", zoneId, err)
		return nil, err
	}
	return zone, nil
}

// enrichAccessPointWithLocation adds location data to a single access point
func enrichAccessPointWithLocation(ap domain.AccessPoint) domain.AccessPoint {
	ap.SetLocation(getAPLocation(ap.GetId()))
	return ap
}

// enrichAccessPointsWithLocation adds location data to a slice of access points
func enrichAccessPointsWithLocation(aps []domain.AccessPoint) []domain.AccessPoint {
	apsWithLocation := make([]domain.AccessPoint, len(aps))
	for i, ap := range aps {
		apsWithLocation[i] = enrichAccessPointWithLocation(ap)
	}
	return apsWithLocation
}

// getAPLocation retrieves the location for an access point
func getAPLocation(apId string) *domain.Location {
	log.Printf("getting location for AP %s", apId)
	location, err := domain.GetLocation(domain.TYPE_AP, apId)
	if err != nil {
		log.Printf("error getting location for AP %s: %v", apId, err)
	}
	return location
}
