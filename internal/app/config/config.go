package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

// ANCHOR: default_config_path

const DefaultConfigPath = "/etc/womblock/config.yaml"

// ANCHOR_END: default_config_path

type Config struct {
	Log Log `yaml:"log"`
}

type Log struct {
	Level  string `validate:"omitempty,oneofci=debug info warn error fatal panic"`
	Format string `validate:"omitempty,oneof=console json"`
}

func Load() (*Config, error) {
	return LoadFromFile(DefaultConfigPath)
}

func LoadFromFile(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config file reading: %v", err)
	}

	dec := yaml.NewDecoder(
		bytes.NewReader(raw),
		yaml.Strict(),
		yaml.DisallowUnknownField(),
	)

	var ret Config
	ret.SetDefaults()

	if err := dec.Decode(&ret); err != nil {
		return nil, fmt.Errorf("config file parsing: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(&ret); err != nil {
		return nil, fmt.Errorf("config file validation: %v", err)
	}

	return &ret, nil
}

func (cfg *Config) SetDefaults() {
	if cfg.Log.Level == "" {
		cfg.Log.Level = "info"
	}
	if cfg.Log.Format == "" {
		cfg.Log.Format = "json"
	}
}
