package routes

import "github.com/gin-gonic/gin"

func PingRoute(router *gin.Engine) {
	// Routes related to zones

	// Ping godoc
	// @Summary Checks if the server is running
	// @Description The GET method is used to check if the server is running.
	// @Id pingGET
	// @Tags ping
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Failure 500
	// @Router /ping [get]
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "pong",
		})
	})
}
