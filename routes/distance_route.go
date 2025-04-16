package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/controllers"
)

func DistanceRoute(router *gin.Engine) {
	router.GET("location/v3/queries/distance", controllers.GetDistance())
}
