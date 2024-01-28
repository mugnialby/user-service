package service

import (
	"log"
	"net/http"

	"github.com/mugnialby/user-service/models/users/entity"
	"github.com/mugnialby/user-service/models/users/request"
	"github.com/mugnialby/user-service/models/users/response"
	"github.com/mugnialby/user-service/repository"
)

func FindByPaging(request *request.FindUserPagingRequest) {

}

func FindById(id int) {

}

func Add(request *request.AddUserRequest) (int, response.Response) {
	user := entity.Users{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		CreatedBy: request.Username,
	}

	result := repository.Add(&user)
	if result.Error != nil {
		log.Printf("failed to add new users : %w", result.Error)
		return http.StatusInternalServerError, response.Response{
			Message: "Failed",
			Data:    nil,
		}
	}

	return http.StatusCreated, response.Response{
		Message: "OK",
		Data:    nil,
	}
}

func Update(request *request.UpdateUserRequest) {

}

func Delete(id int) {

}
