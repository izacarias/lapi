package responses

type UserInfoList struct {
	// ResourceURL Self-referring URL, see note 1.
	ResourceURL string `json:"resourceURL"`
	// User List of users, see note 1.
	User []UserInfo `json:"user"`
}

type UserInfo struct {
	// Address Address of user (e.g. 'sip' URI, 'tel' URI, 'acr' URI) currently on the access point, see note 1.
	Address string `json:"address"`

	// AccessPointId The identity of the access point the user is currently on, see note 1.
	AccessPointId *string `json:"AccessPointId,omitempty"`

	// ZoneId The identity of the zone the user is currently within, see note 1.
	ZoneId string `json:"zoneId"`

	// ResourceURL Self-referring URL, see note 1.
	ResourceURL string `json:"resourceURL"`

	// Date and time that location was collected.
	Timestamp TimeStamp `json:"timestamp"`

	// Location of the User
	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`

	// // CivicInfo Indicates a Civic address
	// CivicInfo            *CivicAddress         `json:"civicInfo,omitempty"`

	// RelativeLocationInfo *RelativeLocationInfo `json:"relativeLocationInfo,omitempty"`

	// AncillaryInfo Reserved for future use.
	AncillaryInfo *string `json:"ancillaryInfo,omitempty"`
}
