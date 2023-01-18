package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password,omitempty" validate:"required"`
}