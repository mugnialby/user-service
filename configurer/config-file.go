package configurer

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ConfigFile() {
	/* CHECK FOR EXTERNAL CONFIG FILES */
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to dev if not specified
	}

	appPath := os.Getenv("APP_PATH")

	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appPath + "/config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("cannot find config file: %w", err))
	}
}
