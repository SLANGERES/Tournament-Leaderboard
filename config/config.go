package config

import (
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type httpServer struct {
	AdminAddress      string `yaml:"admin_address" env-required:"true"`
	UserAddress       string `yaml:"user_address" env-required:"true"`
	TournamentAddress string `yaml:"tournament_addres" env-required:"true"`
}

type Config struct {
	AdminDB    string `yaml:"admin_db"`
	UserDB     string `yaml:"user_db"`
	JwtKey     string `yaml:"jwt_secrate_key"`
	HttpServer httpServer
}

func SetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Unable to get the env ")
	}
	configPath := os.Getenv("ConfigPath")
	if configPath == "" {
		slog.Info("config file is empty ")
	}
	_, err = os.Stat(configPath)

	if os.IsNotExist(err) {
		slog.Info("Unable to find the config file")
	}
	var config Config

	err = cleanenv.ReadConfig(configPath, &config)

	if err != nil {
		slog.Info("Unable to parse the config file")
	}

	return &config
}
