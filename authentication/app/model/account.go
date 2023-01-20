package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type (
	Account struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Email       string             `json:"email,omitempty" bson:"email,omitempty"  validate:"required,email"`
		Password    string             `json:"password,omitempty" bson:"password,omitempty"`
		Phone       string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
		DisplayName string             `json:"displayName,omitempty" bson:"displayName,omitempty" validate:"required"`
		CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
		UpdateAt    time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	}

	AccountSyncSocial struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Email       string             `json:"email,omitempty" bson:"email,omitempty"  validate:"required,email"`
		DisplayName string             `json:"displayName,omitempty" bson:"displayName,omitempty" validate:"required"`
		Social      Social             `json:"social" bson:"social"`
		CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
		UpdateAt    time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
	}
	Social struct {
		Email       string `json:"email,omitempty" bson:"email,omitempty"  validate:"required,email"`
		Provider    string `json:"provider,omitempty" bson:"provider,omitempty" validate:"required"`
		SocialId    string `json:"socialId,omitempty" bson:"socialId,omitempty" validate:"required"`
		AccessToken string `json:"accessToken,omitempty" bson:"accessToken,omitempty" validate:"required"`
	}

	SignInWithEmail struct {
		Email    string `json:"email,omitempty" bson:"email,omitempty"  validate:"required,email"`
		Password string `json:"password,omitempty" bson:"password,omitempty"`
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
