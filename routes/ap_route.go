package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/controllers"
)

func ApRoute(router *gin.Engine) {
	router.GET("location/v3/queries/zones/:id/accessPoints", controllers.ListAccessPoints())
	router.GET("location/v3/queries/zones/:id/accessPoints/:apId", controllers.GetAccessPoint())
}
