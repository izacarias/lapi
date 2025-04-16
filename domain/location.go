package domain

import (
	"errors"
	"time"
)

const (
	TYPE_AP   = "ap"
	TYPE_USER = "user"
)

var (
	ErrLocationNotFound = errors.New("location not found")
)

type Location struct {
	Latitude  float32
	Longitude float32
	Altitude  float32
	// TODO: add the timestamp field
	Timestamp time.Time
}

func NewLocation() *Location {
	return &Location{
		Latitude:  0,
		Longitude: 0,
		Altitude:  0,
		Timestamp: time.Now(),
	}
}

func (l *Location) GetLatitude() float32 {
	return l.Latitude
}

func (l *Location) GetLongitude() float32 {
	return l.Longitude
}

func (l *Location) GetAltitude() float32 {
	return l.Altitude
}

func (l *Location) GetTimestamp() time.Time {
	return l.Timestamp
}
