package auth

import "madi-api/common"

type RegisterWithEmailDto struct {
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	common.BaseSchemas
}

type LoginWithEmailDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
