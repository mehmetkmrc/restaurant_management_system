package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mehmetkmrc/restaurant_management_system"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controller.GetAllInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoiceById())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())
}