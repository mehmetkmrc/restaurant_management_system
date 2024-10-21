package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system/controllers"
) 

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetAllOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrderById())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())

}