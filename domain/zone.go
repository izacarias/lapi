package domain

import (
	"errors"
)

var (
	ErrZoneNotFound = errors.New("zone not found")
)

// Definition of ZoneObject
type Zone struct {
	// Identifier of zone.
	id           string
	accessPoints []AccessPoint
}

func NewZone() *Zone {
	return &Zone{
		id:           "",
		accessPoints: make([]AccessPoint, 0),
	}
}

// SetId : Set the id of the zone
func (z *Zone) SetId(id string) {
	z.id = id
}

func (z *Zone) GetId() string {
	return z.id
}

// add access point to the zone
func (z *Zone) AddAccessPoint(ap AccessPoint) {
	z.accessPoints = append(z.accessPoints, ap)
}

func (z *Zone) GetAccessPoints() []AccessPoint {
	return z.accessPoints
}

func (z *Zone) CountAccessPoints() int {
	return len(z.accessPoints)
}

func (z *Zone) CountSericeableAccessPoints() int {
	count := 0
	for _, ap := range z.accessPoints {
		if ap.GetOperationStatus() == OS_SERVICEABLE {
			count++
		}
	}
	return count
}
