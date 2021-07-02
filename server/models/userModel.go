package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id 						primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname			*string							`json:"firstname" bson:"firstname" validate:"required,min:2,max:100"`
	Lastname			*string							`json:"lastname" bson:"lastname" validate:"required,min:2,max:100"`
	Password			*string							`json:"password" bson:"password" validate:"required,min:6"`
	Email					*string							`json:"email" bson:"email" validate:"required,email"`
	Token					*string							`json:"token" bson:"token"`
	RefreshToken	*string							`json:"refresh_token" bson:"refresh_token"`
	CreatedAt    	time.Time						`json:"created_at" bson:"created_at"`
	UpdatedAt    	time.Time   				`json:"updated_at" bson:"updated_at"`
	UserId    		string   						`json:"user_id" bson:"user_id"`
}
