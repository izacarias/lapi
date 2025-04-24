package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/controllers"
)

func MininetRoute(router *gin.Engine) {
	// Routes related to zones
	router.POST("mininet/user", nil)
	router.POST("mininet/location", controllers.UpdateUserLocation())
	router.POST("mininet/aplocation", controllers.UpdateAPLocation())
	router.POST("mininet/ap", nil)
}
