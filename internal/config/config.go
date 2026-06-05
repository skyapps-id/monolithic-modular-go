package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"-"`
	Modules  []string `yaml:"modules"`
	Preset   string   `yaml:"preset"`
}

type Server struct {
	Addr      string `yaml:"addr"`
	APIPrefix string `yaml:"api_prefix"`
	LogLevel  string `yaml:"log_level"`
}

type Database struct {
	Driver string
	DSN    string
}

const (
	PresetAll       = "all"
	PresetInventory = "inventory"
)

var modulePresets = map[string][]string{
	PresetAll: {
		"users",
		"products",
	},
	PresetInventory: {
		"products",
	},
}

func Load(configPath string) (*Config, error) {
	_ = godotenv.Load()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8080"
	}
	if cfg.Server.APIPrefix == "" {
		cfg.Server.APIPrefix = "/api/v1"
	}

	cfg.Database.Driver = os.Getenv("DB_DRIVER")
	cfg.Database.DSN = os.Getenv("DB_DSN")

	if preset := strings.ToLower(cfg.Preset); preset != "" {
		if modules, ok := modulePresets[preset]; ok {
			cfg.Modules = modules
		}
	}

	return &cfg, nil
}
