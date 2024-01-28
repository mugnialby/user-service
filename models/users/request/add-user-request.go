package request

import (
	"gorm.io/gorm"
)

type AddUserRequest struct {
	gorm.Model
	Username  string `json:"username" validate:"required,max=64"`
	Password  string `json:"password"  validate:"required,max=256"`
	FirstName string `json:"firstname" validate:"required,max=128"`
	LastName  string `json:"lastname" validate:"required,max=128"`
}
