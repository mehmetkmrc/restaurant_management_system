package controllers

import (
	"github.com/gin-gonic/gin"
	
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"go.mongodb.org/mongo-driver/mongo"
)


var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetAllTables() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetTableById() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func CreateTable() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateTable() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}