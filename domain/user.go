package domain

type User struct {
	// AccessPointId The identity of the access point the user is currently on, see note 1.
	ap_id string
	// Address Address of user (e.g. 'sip' URI, 'tel' URI, 'acr' URI) currently on the access point, see note 1.
	address string
	// CivicInfo Indicates a Civic address
	locationInfo string
	tiemstamp    string
	zoneId       string
}

func NewUser(address ...string) *User {
	var addr string
	if len(address) > 0 {
		addr = address[0]
	}

	user := &User{
		ap_id:        "",
		address:      addr,
		locationInfo: "",
		tiemstamp:    "",
		zoneId:       "",
	}

	return user
}
