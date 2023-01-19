package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	RegisterWithEmailDto struct {
		Email       string `json:"email" validate:"required,email"`
		Phone       string `json:"phone" validate:"required"`
		Password    string `json:"password" validate:"required"`
		DisplayName string `json:"displayName" validate:"required"`
	}

	Account struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Email       string             `json:"email,omitempty" bson:"email,omitempty"  validate:"required,email"`
		Password    string             `json:"password,omitempty" bson:"password,omitempty"`
		Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
		DisplayName string             `json:"displayName,omitempty" bson:"displayName,omitempty" validate:"required"`
		CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
		UpdateAt    time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	}
)

func NewAccount(account *Account) *Account {
	return &Account{
		Email:       account.Email,
		Password:    account.Password,
		Phone:       account.Phone,
		DisplayName: account.DisplayName,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
}
