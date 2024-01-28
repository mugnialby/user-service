package response

import "gorm.io/gorm"

type Response struct {
	*gorm.Model
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
