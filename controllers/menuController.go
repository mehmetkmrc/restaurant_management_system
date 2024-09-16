package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"github.com/mehmetkmrc/restaurant_management_system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")


func  GetAllMenus() gin.HandlerFunc{
	return func(c *gin.Context){
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// menus := c.IndentedJSON(http.StatusOK, menus)
		// var menus models.Menu
		// err := menuCollection.Find(ctx.TODO(), bson.M{})
		// defer cancel()
		// if err != nil{
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while inserting the food item"})
		// }
		// c.JSON(http.StatusOK, menus)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing the menu items"})
		}
		var allMenus	[]bson.M
		if err = result.All(ctx, &allMenus); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenuById() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func CreateMenu() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}