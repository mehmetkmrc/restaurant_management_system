package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/mehmetkmrc/restaurant_management_system/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus", controller.GetAllMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenuById())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())
}