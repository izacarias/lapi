package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/controllers"
)

func UserRoute(router *gin.Engine) {
	// Routes related to zones
	router.GET("location/v3/queries/users", controllers.ListUsers())
}
