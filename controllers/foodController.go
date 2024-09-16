package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"github.com/mehmetkmrc/restaurant_management_system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetAllFoods() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetFoodById() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food) 
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return  func(c *gin.Context){

	}
}



func round(num float64) int {
	
}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}