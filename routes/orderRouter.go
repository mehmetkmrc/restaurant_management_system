package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system/controllers"
) 

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetAllOreders())
	incomingRoutes.GET("/orders/:order_id", controller.getOrdersById())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.PATCH("/orders", controller.UpdateOrder())

}