package core

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateInputs(d interface{}) (bool, map[string][]string) {
	validate = validator.New()

	var errors = make(map[string][]string)

	err := validate.Struct(d)

	if err != nil {

		for _, fieldErr := range err.(validator.ValidationErrors) {
			field := strings.ToLower(fieldErr.StructField())
			errors[field] = append(errors[field], fieldErr.Tag())
		}

		return false, errors

	}

	return true, nil
}
