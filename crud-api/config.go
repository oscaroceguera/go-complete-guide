package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

var AppConfig *Config

func LoadAppConfig(){
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("json")    // REQUIRED if the config file does not have the extension in the name

	err := viper.ReadInConfig()    // Find and read the config file
	if(err != nil) {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}