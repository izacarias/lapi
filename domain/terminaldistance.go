package domain

type TerminalDistance struct {
	accuracy  int
	distance  int
	timestamp int
}

// NewTerminalDistance creates a new TerminalDistance instance
func NewTerminalDistance(accuracy, distance, timestamp int) *TerminalDistance {
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
func (td *TerminalDistance) GetTimestamp() int {
	return td.timestamp
}
