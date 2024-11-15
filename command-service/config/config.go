package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type MongoConfig struct {
	Host       string
	Port       string
	Database   string
	Collection string
}

type RedisConfig struct {
	Host string
	Port string
}

type Config struct {
	Mongo MongoConfig
	Redis RedisConfig

	CommandServiceHost string
	CommandServicePort string
}

func Load(path string) (*Config, error) {

	err := godotenv.Load(path + "/.env")
	if err != nil {
		return nil, err
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		Mongo: MongoConfig{
			Host:       conf.GetString("MONGOSH_HOST"),
			Port:       conf.GetString("MONGOSH_PORT"),
			Database:   conf.GetString("MONGOSH_DATABASE"),
			Collection: conf.GetString("MONGOSH_COLLECTION"),
		},
		Redis: RedisConfig{
			Host: conf.GetString("REDIS_HOST"),
			Port: conf.GetString("REDIS_PORT"),
		},

		CommandServiceHost: conf.GetString("COMMANDSERVICE_HOST"),
		CommandServicePort: conf.GetString("COMMANDSERVICE_PORT"),
	}

	if cfg.Mongo.Host == "" || cfg.Mongo.Port == "" || cfg.Mongo.Database == "" || cfg.Mongo.Collection == "" ||
		cfg.Redis.Host == "" || cfg.Redis.Port == "" ||
		cfg.CommandServiceHost == "" || cfg.CommandServicePort == "" {
		return nil, errors.New("configuration values must not be empty")
	}

	return &cfg, nil
}
