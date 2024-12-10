package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/controllers"
)

func ZoneRoute(router *gin.Engine) {
	// Routes related to zones
	router.GET("location/v3/queries/zones", controllers.ListZones())
	router.GET("location/v3/queries/zones/:id", controllers.GetZone())
}
