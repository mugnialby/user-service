package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugnialby/user-service/models/users/request"
	"github.com/mugnialby/user-service/service"
)

func FindByPaging(c *gin.Context) {

	// status, response := service.FindByPaging(&findUserPagingRequest)
	// c.JSON(status, response)
}

func FindById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Add(c *gin.Context) {
	var request request.AddUserRequest
	c.Bind(request)
	valid, message := service.Validate(request)
	if !valid {
		c.JSON(http.StatusBadRequest, message)
	}

	status, response := service.Add(&request)
	c.JSON(status, response)
}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
