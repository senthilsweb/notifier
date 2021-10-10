package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// Setup initialize configuration
func Setup() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./.netlify")
	viper.AddConfigPath("./js")
	viper.AddConfigPath("./functions")
	viper.AddConfigPath("./.netlify/functions")

	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
