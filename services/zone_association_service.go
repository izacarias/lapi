package services

import "github.com/izacarias/lapi/domain/zone_association"

// ZoneAssociationService provides methods to manage zone associations
type ZoneAssociationService struct {
	zoneAssociationRepository zone_association.ZoneAssociationRepository
}

func NewZoneAssociationService(zar zone_association.ZoneAssociationRepository) ZoneAssociationService {
	return ZoneAssociationService{zoneAssociationRepository: zar}
}

func (zas *ZoneAssociationService) GetZone(zoneId string) (zone_association.ZoneAssociation, error) {
	return zas.zoneAssociationRepository.Get(zoneId)
}
