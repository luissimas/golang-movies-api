package config

import (
	"log"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

	if err := load(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}

func load() error {
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		} else {
			log.Print("Config file not found, falling back to default config")
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAPI() APIConfig {
	return cfg.API
}
