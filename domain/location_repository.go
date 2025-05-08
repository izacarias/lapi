package domain

type LocationRepository interface {
	Add(elementType string, elementId string, location *Location) error
	GetLast(elementType string, elementId string) (*Location, error)
}
