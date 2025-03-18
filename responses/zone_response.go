package responses

type ZoneInfo struct {
	// Identifier of zone
	ZoneId string `json:"zoneId"`

	// The number of access points within the zone
	NumberOfAccessPoints int32 `json:"numberOfAccessPoints"`

	// Number of inoperable access points within the zone.
	NumberOfUnserviceableAccessPoints int32 `json:"numberOfUnserviceableAccessPoints"`

	// The number of users currently on the access point.
	NumberOfUsers int32 `json:"numberOfUsers"`

	// Self referring URL
	ResourceURL string `json:"resourceURL"`
}

type ZoneList struct {
	ZoneList struct {
		ResourceURL string     `json:"resourceURL"`
		Zone        []ZoneInfo `json:"zone"`
	} `json:"zoneList"`
}
