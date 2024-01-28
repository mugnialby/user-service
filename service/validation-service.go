package service

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mugnialby/user-service/configurer"
)

func Validate(request interface{}) (bool, string) {
	if request == nil {
		return false, "Request is empty"
	} else {
		if err := configurer.Validator.Struct(request); err != nil {
			if errors, ok := err.(validator.ValidationErrors); ok {
				var message strings.Builder
				for _, e := range errors {
					message.WriteString(strings.ToLower(e.Field()))
					switch e.Tag() {
					case "required":
						message.WriteString(" harus di isi")
					case "max":
						message.WriteString(" tidak boleh lebih dari ")
						message.WriteString(e.Param())
					}

					break
				}

				return false, message.String()
			}
		}

		return true, ""
	}
}
