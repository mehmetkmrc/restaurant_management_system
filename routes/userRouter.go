package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.POST("/users/signup", controller.SignUp())
}