package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Menu struct{
	ID primitive.ObjectID `bson:"_id"`
	Name				  string 	`json:"name" required:"validate"`
	Category 			  string   	`json:"category" required:"validate"`
	Start_Date			  *time.Time `json:""`
	End_Date			  *time.Time
	Created_at			  time.Time
	Updated_at			  time.Time
	Menu_id				  string
}