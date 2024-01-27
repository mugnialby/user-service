package request

import (
	"gorm.io/gorm"
)

type FindUserPagingRequest struct {
	gorm.Model
	Username  string `json:"username" validate:"required, max=64"`
	FirstName string `json:"firstname" validate:"required, max=128"`
	LastName  string `json:"lastname" validate:"required, max=128"`
	Page      int    `json:"page" validate:"required, max=6"`
	PageSize  int    `json:"pagesize" validate:"required, max=2"`
	SortOrder string `json:"sortorder" validate:"required, max=3"`
}
