package models

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
	AccessPointId string `bson:"id"`

	// LocationInfo LocationInfo `bson:"locationInfo,omitempty,inline"`

	ConnectionType ConnectionType `bson:"connection_type"`

	OperationStatus OperationStatus `bson:"operation_status"`

	// Number of users currently on the access point.
	NumberOfUsers int32 `bson:"number_users"`

	// Time zone of access point.
	Timezone string `bson:"timezone,omitempty"`

	// Interest realm of access point.
	// InterestRealm string `json:"interestRealm,omitempty"`
}
