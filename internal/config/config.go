package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT"`
		Host string `yaml:"host" env:"SERVER_HOST"`
	} `yaml:"server"`

	Logging struct {
		Level  string `yaml:"level" env:"LOG_LEVEL"`
		Format string `yaml:"format" env:"LOG_FORMAT"`
	} `yaml:"logging"`

	Gin struct {
		Mode string `yaml:"mode" env:"GIN_MODE"`
	} `yaml:"gin"`

	App struct {
		Name    string `yaml:"name" env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	} `yaml:"app"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := &Config{}

	if err := loadFromYAML(cfg); err != nil {
		log.Printf("Warning: Could not load config from YAML: %v", err)
	}

	loadFromEnv(cfg)

	setDefaults(cfg)

	return cfg
}

func loadFromYAML(cfg *Config) error {
	data, err := os.ReadFile("./internal/handlers/config/config.yml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}

func loadFromEnv(cfg *Config) {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		cfg.Server.Port = port
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Host = host
	}

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		cfg.Logging.Level = level
	}

	if mode := os.Getenv("GIN_MODE"); mode != "" {
		cfg.Gin.Mode = mode
	}
}

func setDefaults(cfg *Config) {
    if cfg.Server.Port == "" {
        cfg.Server.Port = "8080"
    }
    if cfg.Server.Host == "" {
        cfg.Server.Host = "localhost"
    }
    if cfg.Logging.Level == "" {
        cfg.Logging.Level = "info"
    }
    if cfg.Gin.Mode == "" {
        cfg.Gin.Mode = "debug"
    }
    if cfg.App.Name == "" {
        cfg.App.Name = "Task Manager API"
    }
}