package responses

// TimeStamp represents a timestamp with seconds and nanoseconds
type TimeStamp struct {
	// The seconds part of the time. Time is defined as Unix-time since January 1, 1970, 00:00:00 UTC.
	Seconds uint32 `json:"seconds"`

	// The nanoseconds part of the time.
	NanoSeconds uint32 `json:"nanoSeconds"`
}
