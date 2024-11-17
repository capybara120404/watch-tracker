package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogFile          string
	NameOFLogger     string
	ConnectionString string
}

func NewConfig(configFile string) (*Config, error) {
	err := godotenv.Load(configFile)
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	return &Config{
		LogFile:          os.Getenv("LOG_FILE"),
		NameOFLogger:     os.Getenv("NAME_OF_LOGGER"),
		ConnectionString: os.Getenv("CONNECTION_STRING"),
	}, nil
}
