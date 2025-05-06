package routes

import (
	"github.com/gin-gonic/gin"
	// Swagger docs
	"github.com/izacarias/lapi/configs"
	_ "github.com/izacarias/lapi/docs" // This line is necessary for go-swagger to find your docs!
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoute(router *gin.Engine) {
	// Routes related to Swagger
	rootUrl := configs.GetConfigApiRoot()
	url := ginSwagger.URL(rootUrl + "/swagger/doc.json")
	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
