package zone_association

import "errors"

var (
	ErrZoneNotFound = errors.New("zone not found")
	ErrFailedToAdd  = errors.New("failed to add zone association")
)

type ZoneAssociationRepository interface {
	Get(zoneId string) (ZoneAssociation, error)
	GetAll() ([]ZoneAssociation, error)
	Add(ZoneAssociation) error
}
