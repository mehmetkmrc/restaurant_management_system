package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"github.com/mehmetkmrc/restaurant_management_system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type InvoiceViewFormat struct {
	Invoice_id			string
	Payment_method		string
	Order_id			string
	Payment_status		*string
	Payment_due			interface{}
	Table_number		interface{}
	Payment_due_date	time.Time
	Order_details		interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetAllInvoices() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing order items"})
	}

	var allInvoices []bson.M
	if err = result.All(ctx, &allInvoices); err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, allInvoices)
	}

	
}

func GetInvoiceById() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")

		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id":invoiceId}).Decode(&invoice)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while listing invoice item"})
		}
		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil{
			invoiceView.Payment_method = *invoice.Payment_method
		}
	}
}

func CreateInvoice() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateInvoice() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}