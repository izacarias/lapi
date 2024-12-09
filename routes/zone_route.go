package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/controllers"
)

func ZoneRoute(router *gin.Engine) {
	// Routes related to zones
	zoneQueryURI := configs.GetConfigQueriesURI()
	router.GET(zoneQueryURI+"/zones", controllers.ListZones())
	router.GET(zoneQueryURI+"/zones/:id", controllers.GetZone())
}
