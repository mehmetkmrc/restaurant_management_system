package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"github.com/mehmetkmrc/restaurant_management_system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

var validate = validator.New()

func GetAllFoods() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi("recordPerPage") //Bu fonksiyon verilen bir string ifadesini integer(tam sayı) türüne dönüştürmek için kullanılır.
		if err != nil || recordPerPage < 1{
			recordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err !=nil || page < 1 {
			page = 1
		}

		startIndex := (page-1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}} }}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
				}
			}
		}
		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage
		})
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured"})
		}
		var GetAllFoods []bson.M
		if err = result.All(ctx, &GetAllFoods); err !=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, GetAllFoods[0])
	}
}

func GetFoodById() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")//Burada Param metodunun kullanılmasının sebebi food_id url de dinamik bir parametreye karşılık geliyor
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food) 
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the food item"})
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return  func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)		
		var menu models.Menu //The writer added
		var food models.Food
		if err := c.Bind(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return  
		}
		err := menuCollection.FindOne(ctx, bson.M{"new_id":food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil{
			 msg := fmt.Sprintf("menu not found")
			 c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			 return 
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil{
			msg := 	fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
func round(num float64) int {
	
}

func toFixed(num float64, precision int) float64 {

}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}