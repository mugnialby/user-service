package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func DbConnection() {
	host := "host=" + viper.GetString("database.host") + " "
	user := "user=" + viper.GetString("database.user") + " "
	password := "password=" + viper.GetString("database.password") + " "
	dbname := "dbname=" + viper.GetString("database.dbname") + " "
	port := "port=" + viper.GetString("database.port") + " "
	sslMode := "sslmode=" + viper.GetString("database.sslmode") + " "
	timeZone := "TimeZone=" + viper.GetString("database.timeZone")
	dsn := host + user + password + dbname + port + sslMode + timeZone

	Connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("cannot connect to database: %w", err))
	}

	database, e := Connection.DB()
	if err != nil {
		panic(fmt.Errorf("cannot initialize database: %w", e))
	}

	database.SetMaxIdleConns(viper.GetInt("database.maxIdleConnection"))
	database.SetMaxOpenConns(viper.GetInt("database.maxOpenConnection"))
	defer database.Close()
}
