package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
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
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}
)

// Parse config from yaml file.
func Parse() (*Config, error) {

	path := fetchConfigPath()

	cfg := Config{}
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal cofig from yaml file: %w", err)
	}
	return &cfg, nil
}

func fetchConfigPath() string {
	res, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		panic("could not find config file")
	}
	return res
}

// init in config is for loading env variables from env file
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
