package utils

import (
	"fmt"
	"net/http"

	"github.com/izacarias/lapi/configs"
)

func GetConfigResourceURL() string {
	resourceURL := fmt.Sprintf("%s/location/%s", configs.GetConfigApiRoot(), configs.GetConfigApiVersion())
	return resourceURL
}

func GetConfigQueriesURI() string {
	return fmt.Sprintf("%s/queries", GetConfigResourceURL())
}

func GetConfigQueriesZones() string {
	return fmt.Sprintf("%s/zones", GetConfigQueriesURI())
}

func ConstructZoneResourceUrl(r *http.Request, zoneId string) string {
	zonesURI := GetConfigQueriesZones()
	return fmt.Sprintf("%s/%s", zonesURI, zoneId)
}

func ConstructZoneListResourceUrl(r *http.Request) string {
	zonesURI := GetConfigQueriesZones()
	return fmt.Sprintf("%s", zonesURI)
}
