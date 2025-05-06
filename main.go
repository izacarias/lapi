package main

import (
	"github.com/gin-gonic/gin"
	"github.com/izacarias/lapi/configs"
	"github.com/izacarias/lapi/mock"
	"github.com/izacarias/lapi/routes"
)

// @title ETSI GS MEC 013 - Location API
// @version 3.1.1
// @description The ETSI MEC ISG MEC013 Location API described using OpenAPI

// @contact.name Iulisloi Zacarias
// @contact.url https://github.com/izacarias
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs.ETSI MEC013 V3.1.1 Location API
// @externalDocs.https://www.etsi.org/deliver/etsi_gs/MEC/001_099/013/03.01.01_60/gs_mec013v030101p.pdf
// @tags.name location
// @schemes http
// @host localhost:8080
// @basePath /location/v3/
func main() {
	// Connect to MongoDB
	client := configs.ConnectDB()

	// set the gin mode to release
	// gin.SetMode(gin.ReleaseMode)

	// Inserting mock data
	mock.InsertMockData(client)

	router := gin.Default()

	// Setting up routes
	routes.PingRoute(router)
	routes.ZoneRoute(router)
	routes.ApRoute(router)
	routes.UserRoute(router)
	routes.DistanceRoute(router)
	routes.MininetRoute(router)

	// Adding Swagger routes
	routes.SwaggerRoute(router)

	// Running the server
	router.Run(":8080")
}
