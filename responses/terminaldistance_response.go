package responses

type TerminalDistance struct {
	// Accuracy Accuracy of the provided distance in meters
	Accuracy int `json:"accuracy,omitempty"`

	// Distance Distance from terminal to a location or between two terminals specified in meters
	Distance  int        `json:"distance"`
	Timestamp *TimeStamp `json:"timestamp,omitempty"`
}
