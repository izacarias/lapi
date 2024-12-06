package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/izacarias/lapi/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var zoneCollection *mongo.Collection = configs.GetCollection(configs.DB, "zones")
var validate = validator.New()

func CreateZone() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// var zone models.Zone
		// defer cancel()

		// if err := c.ShouldBindJSON(&zone); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// if err := validate.Struct(zone); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		// result, err := zoneCollection.InsertOne(context.Background(), zone)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating zone"})
		// 	return
		// }
		// c.JSON(http.StatusCreated, result)
	}
}
