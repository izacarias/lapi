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

	// Adding Ping routes
	routes.PingRoute(router)
	// Adding Zone routes
	routes.ZoneRoute(router)

	// Adding Swagger routes
	routes.SwaggerRoute(router)

	// Running the server
	router.Run(":8080")
}
