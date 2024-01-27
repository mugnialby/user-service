package main

import (
	"github.com/mugnialby/user-service/configurer"
	"github.com/mugnialby/user-service/database"
	"github.com/mugnialby/user-service/models/users/entity"
)

func init() {
	configurer.GetConfigFile()
	database.DbConnection()
}

func main() {
	database.DB.AutoMigrate(
		&entity.Users{},
	)
}
