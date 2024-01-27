package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mugnialby/user-service/configurer"
	"github.com/mugnialby/user-service/controller"
	"github.com/mugnialby/user-service/database"
	"github.com/spf13/viper"
)

func init() {
	configurer.ConfigFile()
	configurer.DataValidation()
	database.DbConnection()
}

func main() {
	router := gin.Default()

	router.GET("/", controller.FindByPaging)
	router.GET("/find", controller.FindById)
	router.POST("/add", controller.Add)
	router.PATCH("/update", controller.Update)
	router.GET("/delete", controller.Delete)

	router.Run(":" + viper.GetString("app.port"))
}
