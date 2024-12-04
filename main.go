package main

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/mock"
	"github.com/izacarias/lapi/routes"
)

func main() {
	// Connect to MongoDB
	client := configs.ConnectDB()

	// Inserting mock data
	mock.InsertMockData(client)

	router := gin.Default()

	// Adding routes
	routes.ZoneRoute(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.Run(":8080")
}
