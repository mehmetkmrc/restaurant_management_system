package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mehmetkmrc/restaurant_management_system/database"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")


/*Burada GetUsers, GetUser, Login ve SignUp fonksiyonlarının return func(c *gin.Context) şeklinde bir return değer almasının sebebi Gin Framework'ünde HTTP isteklerini alıp ve cevap vermesi. Yani HTTP ile etkileşimde bulunurlar.
*/
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}


/*
Buradaki HashPassword ve VerifyPassword fonksiyonları ise HTTP isteklerini işlemek için değil, uygulama içinde kullanılmak üzere yardımcı (utility) fonksiyonlardır.
*/

func HashPassword(password string) string{

}


func VerifyPassword(userPassword string, providePassword string)(bool, string){

}