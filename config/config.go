package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type httpServer struct {
	AdminAddress      string `yaml:"admin_address" env-required:"true"`
	UserAddress       string `yaml:"user_address" env-required:"true"`
	TournamentAddress string `yaml:"tournament_address" env-required:"true"`
}

type Config struct {
	AdminDB    string     `yaml:"admin_db"`
	UserDB     string     `yaml:"user_db"`
	JwtKey     string     `yaml:"jwt_secrate_key"`
	HttpServer httpServer `yaml:"http_server"`
}

// SetConfig loads environment variables and reads the application configuration
// from the file specified in the ConfigPath environment variable.


func SetConfig() (*Config, error) {
	if err := godotenv.Load("config/.env"); err != nil {
		slog.Warn("Unable to load .env file", slog.Any("error", err))
	}

	configPath := os.Getenv("ConfigPath")
	if configPath == "" {
		return nil, fmt.Errorf("ConfigPath environment variable is empty")
	}

	info, err := os.Stat(configPath)
	if err != nil {
		return nil, fmt.Errorf("unable to access config file: %w", err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("config path is a directory, expected a file: %s", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		slog.Warn("failed to parse config file", slog.Any("error", err))
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}
