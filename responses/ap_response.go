package responses

// List of ConnectionType
const (
	CT_LTE     ConnectionType = "LTE"
	CT_WIFI    ConnectionType = "Wifi"
	CT_WIMAX   ConnectionType = "Wimax"
	CT_NR5G    ConnectionType = "5G NR"
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

type AccessPointInfo struct {

	// Identifier of access point.
	AccessPointId string `json:"accessPointId"`

	// LocationInfo LocationInfo `json:"locationInfo,omitempty"`

	ConnectionType ConnectionType `json:"connectionType"`

	OperationStatus OperationStatus `json:"operationStatus"`

	// Number of users currently on the access point.
	NumberOfUsers int32 `json:"numberOfUsers"`

	// Time zone of access point.
	Timezone string `json:"timezone,omitempty"`

	// Interest realm of access point.
	InterestRealm string `json:"interestRealm,omitempty"`

	// Self referring URL
	ResourceURL string `json:"resourceURL"`
}

type AnAccessPointInfo struct {
	AccessPoint AccessPointInfo `json:"accessPointInfo"`
}

// AccessPointList - A type containing list of access points.
type AccessPointInfoList struct {
	// Identifier of zone
	ZoneId string `json:"zoneId"`

	// Collection of the access point information list.
	AccessPoint []AccessPointInfo `json:"accessPoint,omitempty"`

	// Self referring URL
	ResourceURL string `json:"resourceURL"`
}

type AccessPointList struct {
	AccessPoint AccessPointInfoList `json:"accessPointList"`
}
