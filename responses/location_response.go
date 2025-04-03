package responses

// Defines values for LocationInfoShape.
const (
	LocationInfoShapeN1 LocationInfoShape = 1
	LocationInfoShapeN2 LocationInfoShape = 2
	LocationInfoShapeN3 LocationInfoShape = 3
	LocationInfoShapeN4 LocationInfoShape = 4
	LocationInfoShapeN5 LocationInfoShape = 5
	LocationInfoShapeN6 LocationInfoShape = 6
	LocationInfoShapeN7 LocationInfoShape = 7
)

type LocationInfoShape int

type LocationInfo struct {

	// Location altitude relative to the WGS84 ellipsoid surface.
	Altitude float32 `json:"altitude,omitempty"`

	// Location latitude, expressed in the range -90° to +90°. Cardinality greater than one only if \"shape\" equals 7.
	Latitude []float32 `json:"latitude"`

	// Location longitude, expressed in the range -180° to +180°. Cardinality greater than one only if \"shape\" equals 7.
	Longitude []float32 `json:"longitude"`

	/* Shape information, as detailed in [14], associated with the reported location coordinate:
	1 = Ellipsoid_Arc
	2 = ellipsoid_Point
	3 = ellipsoid_Point_Altitude
	4 = ellipsoid_Point_Altitude_Uncert_Ellipsoid
	5 = ellipsoid_Point_Uncert_Circle
	6 = ellipsoid_Point_Uncert_Ellipse
	7 = polygon */
	Shape LocationInfoShape `json:"shape"`

	// Present only if \"shape\" equals 6.
	// UncertaintyRadius int32 `json:"uncertaintyRadius,omitempty"`

	// Present only if \"shape\" equals 6.
	// InnerRadius int32 `json:"innerRadius,omitempty"`

	// Present only if \"shape\" equals 6.
	// IncludedAngle int32 `json:"includedAngle,omitempty"`

	// Confidence by which the position of a target entity is known to be within the shape description, expressed as a percentage and defined in [14].
	// Present only if \"shape\" equals 1, 4 or 6.
	// Confidence int32 `json:"confidence,omitempty"`

	// Horizontal accuracy/(semi-major) uncertainty of location provided in meters, as defined in [14].
	// Present only if \"shape\" equals 4, 5 or 6.
	// Accuracy int32 `json:"accuracy,omitempty"`

	// Altitude accuracy/uncertainty of location provided in meters, as defined in [14].
	// Present only if \"shape\" equals 3 or 4.
	// AccuracyAltitude int32 `json:"accuracyAltitude,omitempty"`

	// Horizontal accuracy/(semi-major) uncertainty of location provided in meters, as defined in [14].
	// Present only if \"shape\" equals 4, 5 or 6.
	// AccuracySemiMinor int32 `json:"accuracySemiMinor,omitempty"`

	// Present only if \"shape\" equals 6.
	// OffsetAngle int32 `json:"offsetAngle,omitempty"`

	// Angle of orientation of the major axis, expressed in the range 0° to 180°, as defined in [14].
	// Present only if \"shape\" equals 4 or 6.
	// OrientationMajorAxis int32 `json:"orientationMajorAxis,omitempty"`

	// Velocity Velocity `json:"velocity,omitempty"`
}
