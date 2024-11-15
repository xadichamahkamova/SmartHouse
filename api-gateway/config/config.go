package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


type Config struct {
	UserServiceHost string
	UserServicePort string

	DeviceServiceHost string
	DeviceServicePort string 

	CommandServiceHost string
	CommandServicePort string

	ApiGatewayHost string
	ApiGatewayPort string
}

func Load(path string) (*Config, error) {

	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, err
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		UserServiceHost: conf.GetString("USERSERVICE_HOST"),
		UserServicePort: conf.GetString("USERSERVICE_PORT"),

		DeviceServiceHost: conf.GetString("DEVICESERVICE_HOST"),
		DeviceServicePort: conf.GetString("DEVICESERVICE_PORT"),

		CommandServiceHost: conf.GetString("COMMANDSERVICE_HOST"),
		CommandServicePort: conf.GetString("COMMANDSERVICE_PORT"),

		ApiGatewayHost: conf.GetString("APIGATEWAY_HOST"),
		ApiGatewayPort: conf.GetString("APIGATEWAY_PORT"),
	}
	if cfg.UserServiceHost == "" || cfg.UserServicePort == "" ||
		cfg.DeviceServiceHost == "" || cfg.DeviceServicePort == "" ||
		cfg.CommandServiceHost == "" || cfg.CommandServicePort == "" ||
		cfg.ApiGatewayHost == "" || cfg.ApiGatewayPort == "" {
		return nil, errors.New("configuration values must not be empty")
	}

	return &cfg, nil
}
