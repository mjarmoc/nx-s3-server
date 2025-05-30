package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("default")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing default configuration file")
	}

	envConfig := viper.New()
	envConfig.SetConfigType("yaml")
	envConfig.AddConfigPath("config/")
	envConfig.SetConfigName(env)
	err = envConfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing env configuration file")
	}

	config.MergeConfigMap(envConfig.AllSettings())
}

func GetConfig() *viper.Viper {
	return config
}
