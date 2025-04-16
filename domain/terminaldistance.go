package domain

import "time"

type TerminalDistance struct {
	accuracy  int
	distance  int
	timestamp time.Time
}

// NewTerminalDistance creates a new TerminalDistance instance
func NewTerminalDistance(accuracy, distance int, timestamp time.Time) *TerminalDistance {
	return &TerminalDistance{
		accuracy:  accuracy,
		distance:  distance,
		timestamp: timestamp,
	}
}

// GetAccuracy returns the accuracy of the distance
func (td *TerminalDistance) GetAccuracy() int {
	return td.accuracy
}

// Get Distance returns the distance
func (td *TerminalDistance) GetDistance() int {
	return td.distance
}

// GetTimestamp returns the timestamp
func (td *TerminalDistance) GetTimestamp() time.Time {
	return td.timestamp
}
