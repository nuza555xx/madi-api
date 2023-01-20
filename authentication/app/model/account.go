package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type (
	Account struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Email       string             `json:"email" bson:"email" validate:"required,email"`
		Password    string             `json:"password" bson:"password" validate:"required,gte=8"`
		Phone       *Phone             `json:"phone" bson:"phone" validate:"required"`
		DisplayName string             `json:"displayName" bson:"displayName" validate:"required"`
		CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
		UpdateAt    time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	}

	AccountSyncSocial struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Email       string             `json:"email" bson:"email" validate:"required,email"`
		DisplayName string             `json:"displayName" bson:"displayName" validate:"required"`
		Social      *Social            `json:"social" bson:"social" validate:"required"`
		CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
		UpdateAt    time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	}
	Social struct {
		Email       string `json:"email" bson:"email" validate:"required,email"`
		Provider    string `json:"provider" bson:"provider" validate:"required"`
		SocialId    string `json:"socialId" bson:"socialId" validate:"required"`
		AccessToken string `json:"accessToken" bson:"accessToken" validate:"required"`
	}

	Phone struct {
		Tel         string `json:"tel," bson:"tel," validate:"required"`
		CountryCode string `json:"countryCode," bson:"countryCode," validate:"required"`
	}

	SignInWithEmail struct {
		Email    string `json:"email" bson:"email" validate:"required,email"`
		Password string `json:"password" bson:"password" validate:"required"`
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

func NewAccountSyncSocial(account *AccountSyncSocial) *AccountSyncSocial {
	return &AccountSyncSocial{
		Email:       account.Email,
		DisplayName: account.DisplayName,
		Social:      account.Social,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
}

func (account *Account) GenerateFromPassword(password string) error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	account.Password = string(newPassword)
	return nil
}
