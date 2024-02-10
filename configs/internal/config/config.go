package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Env      string        `yaml:"env" env-default:"local"`
		DB       DB            `yaml:"db"`
		TokenTTL time.Duration `yaml:"token_ttl"`
		GRPC     GRPCConfig    `yaml:"grpc"`
	}

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	GRPCConfig struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	}
)

// Parse config from yaml file.
func Parse(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config from yaml file: %w", err)
	}
	cfg := Config{}
	err = yaml.Unmarshal(data)
}
