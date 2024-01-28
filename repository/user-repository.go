package repository

import (
	"github.com/mugnialby/user-service/database"
	"github.com/mugnialby/user-service/models/users/entity"
	"github.com/mugnialby/user-service/models/users/request"
	"gorm.io/gorm"
)

func FindByPaging(findByPagingRequest *request.FindUserPagingRequest) {
	// database.Connection.Find()
}

func FindById(id *int) {
	// database.Connection.Find()
}

func Add(user *entity.Users) *gorm.DB {
	return database.Connection.Create(user)
}

func Update(updateUserRequest *request.UpdateUserRequest) {
	// database.Connection.Find()
}

func Delete(id *int) {
	// database.Connection.Find()
}
