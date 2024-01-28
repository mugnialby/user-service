package entity

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint64
	Username  string
	Password  string
	FirstName string
	LastName  string
	CreatedBy string
	// CreatedAt date
	ModifiedBy string
	// ModifiedAt date
}
