package response

import "gorm.io/gorm"

type Response struct {
	*gorm.Model
	Message string
	Data    interface{}
}
