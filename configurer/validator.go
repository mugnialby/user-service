package configurer

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

func DataValidation() {
	if Validator == nil {
		Validator = validator.New()
	}
}
