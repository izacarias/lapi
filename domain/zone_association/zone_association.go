// Zone Association is an aggregate to hold the relationship between a zone and its access points.
package zone_association

import (
	"errors"

	"github.com/izacarias/lapi/domain/access_point"
	"github.com/izacarias/lapi/domain/zone"
)

var (
	ErrZoneIdEmpty = errors.New("zone id cannot be empty")
)

type ZoneAssociation struct {
	zone          zone.Zone
	access_points []access_point.AccessPoint
}

func NewZoneAssociation(zoneId string) (ZoneAssociation, error) {
	if zoneId == "" {
		return ZoneAssociation{}, ErrZoneIdEmpty
	}

	z := zone.Zone{Id: zoneId}
	za := ZoneAssociation{zone: z}
	return za, nil
}

func (za *ZoneAssociation) GetZoneId() string {
	return za.zone.Id
}
