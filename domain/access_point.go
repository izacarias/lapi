package domain

import "errors"

// List of ConnectionType
const (
	CT_LTE     ConnectionType = "LTE"
	CT_WIFI    ConnectionType = "Wifi"
	CT_WIMAX   ConnectionType = "Wimax"
	CT_5GNR    ConnectionType = "5G NR"
	CT_UNKNOWN ConnectionType = "UNKNOWN"
)

// List of OperationStatus
const (
	OS_SERVICEABLE   OperationStatus = "Serviceable"
	OS_UNSERVICEABLE OperationStatus = "Unserviceable"
	OS_UNKNOWN       OperationStatus = "UNKNOWN"
)

// ConnectionType : This enumeration represents the connection type of an access point
type ConnectionType string

// OperationStatus : An enumeration defining the operations status of an access point.
type OperationStatus string

type AccessPoint struct {
	// Identifier of access point.
	id               string
	connection_type  ConnectionType
	operation_status OperationStatus
	time_zone        string
	// list of users connected to the AP
	users    []User
	location Location
	zone_id  string
}

var (
	ErrAccessPointNotFound = errors.New("access point not found")
)

func NewAccessPoint() *AccessPoint {
	return &AccessPoint{
		id:               "",
		connection_type:  CT_UNKNOWN,
		operation_status: OS_UNKNOWN,
		time_zone:        "",
		location:         *NewLocation(),
	}
}

// SetId : Set the id of the access point
func (ap *AccessPoint) SetId(id string) {
	ap.id = id
}

func (ap *AccessPoint) GetId() string {
	return ap.id
}

// SetConnectionType : Set the connection type of the access point
func (ap *AccessPoint) SetConnectionType(ct ConnectionType) {
	ap.connection_type = ct
}

func (ap *AccessPoint) GetConnectionType() ConnectionType {
	return ap.connection_type
}

// SetOperationStatus : Set the operation status of the access point
func (ap *AccessPoint) SetOperationStatus(os OperationStatus) {
	ap.operation_status = os
}

func (ap *AccessPoint) GetOperationStatus() OperationStatus {
	return ap.operation_status
}

// SetTimeZone : Set the time zone of the access point
func (ap *AccessPoint) SetTimeZone(tz string) {
	ap.time_zone = tz
}

func (ap *AccessPoint) GetTimeZone() string {
	return ap.time_zone
}

// AddUser : Add a user to the access point
func (ap *AccessPoint) AddUser(user *User) {
	ap.users = append(ap.users, *user)
}

func (ap *AccessPoint) CountUsers() int {
	return len(ap.users)
}

func (ap *AccessPoint) SetLocation(location *Location) {
	ap.location = *location
}

func (ap *AccessPoint) GetLocation() Location {
	return ap.location
}

func (ap *AccessPoint) SetZoneId(zoneId string) {
	ap.zone_id = zoneId
}

func (ap *AccessPoint) GetZoneId() string {
	return ap.zone_id
}
