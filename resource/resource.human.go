package resource

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Human struct {
	Name string `json:"user_name" bson:"user_name" binding:"required"`
	CreatedAT time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DbHuman struct {
	ID primitive.ObjectID `json:"user_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"user_name,omitempty" bson:"user_name,omitempty" binding:"required"`
	CreatedAT time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}