package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	defaultMinKindnessLevel int = 6
)

type Config struct {
	Application struct {
		MinKindnessLevel int `yaml:"min-herbo-kindness"`
	} `yaml:"application"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, err
}

func (c *Config) Validate() error {
	if c.Application.MinKindnessLevel < 1 || c.Application.MinKindnessLevel > 11 {
		return ErrInvalidConfig
	}
	return nil
}

func (c *Config) SetDefaults() {
	if c.Application.MinKindnessLevel == 0 {
		c.Application.MinKindnessLevel = defaultMinKindnessLevel
	}
}
