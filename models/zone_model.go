package models

type Zone struct {
	Id               string   `bson:"id"`
	AccessPointsList []string `bson:"access_points_list"`
}

// Get the number of access points in a zone
func (z *Zone) CoungAccessPoints() int {
	return len(z.AccessPointsList)
}

// Get the number of serviceable access points in a zone
func (z *Zone) C() int {
	// for _, apId := range z.AccessPointsList {
	// 	// ap, err := GetAccessPointById(apId)
	// 	// if ap.OperationStatus == OS_SERVICEABLE {
	// 	// 	count++
	// 	// }
	// }
	panic("not implemented")
}
