package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"github.com/mehmetkmrc/restaurant_management_system/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct{
	Table_id *string
	Order_items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetAllOrderItems() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetOrderItemById() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error){
	
}

func CreateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}