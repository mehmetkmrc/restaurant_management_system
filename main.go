package main

import (
	"os"
	middleware "github.com/mehmetkmrc/restaurant_management_system/middleware"
	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/routes"
)



func main(){
	port := os.Getenv("PORT")

	if port  == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}