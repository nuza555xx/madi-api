package core

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateInputs(d interface{}) (bool, map[string][]string) {
	validate = validator.New()

	err := validate.Struct(d)

	if err != nil {

		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		errors := make(map[string][]string)

		for _, err := range err.(validator.ValidationErrors) {

			name := strings.ToLower(err.StructField())
			errors[name] = append(errors[name], err.Tag())
			// switch err.Tag() {
			// case "required":
			// 	errors[name] = append(errors[name], "The "+name+" is required.")
			// 	break
			// case "alpha":
			// 	errors[name] = append(errors[name], "The "+name+" should contain only letters.")
			// 	break
			// case "eqfield":
			// 	errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
			// 	break
			// case "gte":
			// 	errors[name] = append(errors[name], "The "+name+" should greater than "+err.Param()+" characters.")
			// 	break
			// case "lte":
			// 	errors[name] = append(errors[name], "The "+name+" should less than "+err.Param()+" characters.")
			// 	break
			// default:
			// 	errors[name] = append(errors[name], "The "+name+" is invalid.")
			// 	break
			// }
		}

		return false, errors
	}
	return true, nil
}
