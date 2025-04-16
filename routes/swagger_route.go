package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/izacarias/lapi/docs" // This line is necessary for go-swagger to find your docs!
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoute(router *gin.Engine) {
	// Routes related to Swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
