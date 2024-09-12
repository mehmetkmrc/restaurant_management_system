package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controller.GetAllFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFoodById())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}