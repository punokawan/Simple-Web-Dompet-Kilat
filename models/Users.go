package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" binding:"required"`
}
