package domain

type User struct {
	Address     string
	AccessPoint string
	ZoneId      string
	Location    *Location
}

// GetAddress returns the user's address
func (u *User) GetAddress() string {
	return u.Address
}

// SetAddress sets the user's address
func (u *User) SetAddress(address string) {
	u.Address = address
}

// GetAccessPoint returns the user's access point ID
func (u *User) GetAccessPoint() string {
	return u.AccessPoint
}

// SetAccessPoint sets the user's access point ID
func (u *User) SetAccessPoint(accessPoint string) {
	u.AccessPoint = accessPoint
}

// GetZoneId returns the user's zone ID
func (u *User) GetZoneId() string {
	return u.ZoneId
}

// SetZoneId sets the user's zone ID
func (u *User) SetZoneId(zoneId string) {
	u.ZoneId = zoneId
}

// GetLocation returns the user's location
func (u *User) GetLocation() *Location {
	return u.Location
}

// SetLocation sets the user's location
func (u *User) SetLocation(location *Location) {
	u.Location = location
}
