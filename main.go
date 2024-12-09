package main

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/mock"
	"github.com/izacarias/lapi/routes"
)

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /location/v3/
// @schemes http
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
