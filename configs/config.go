package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Configs Configuration

type Server struct {
	Host string
	Port string
}

type Connection struct {
	Max int
}

type Storage struct {
	Path string
}

type Configuration struct {
	Server     Server
	Connection Connection
	Storage    Storage
}

var path = "./configs"

func setConfigPath(cPath string) {
	path = cPath
}

func GetConfigs() *Configuration {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(&Configs)
	return &Configs
}
