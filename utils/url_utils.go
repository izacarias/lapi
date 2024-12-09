package utils

import (
	"fmt"
	"net/http"

	"github.com/izacarias/lapi/configs"
)

var (
	API_VERSION = "v1"
	API_PATH    = "api"
)

func ConstructZoneResourceUrl(r *http.Request, zoneId string) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := r.Host
	zonesURI := configs.GetConfigQueriesZones()
	return fmt.Sprintf("%s://%s/%s/%s", scheme, host, zonesURI, zoneId)
}
