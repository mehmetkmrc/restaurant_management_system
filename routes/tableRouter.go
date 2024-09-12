package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", controller.GetAllTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTableById())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id" , controller.UpdateTable())
}