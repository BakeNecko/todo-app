package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App     `yaml:"app"`
		Auth    `yaml:"auth"`
		Secret  `yaml:"secret"`
		HTTP    `yaml:"http"`
		Log     `yaml:"logger"`
		MongoDB `yaml:"db"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	Auth struct {
		AccessTokenTTL  string `env-required:"true" yaml:"accessTokenTTL"    env:"ACCESS_TOKEN_TTL"`
		RefreshTokenTTL string `env-required:"true" yaml:"refreshTokenTTL"    env:"REFRESH_TOKEN_TTL"`
	}
	Secret struct {
		Key  string `env-required:"true" yaml:"key"                env:"SECRET_KEY"`
		Salt string `env-required:"true" yaml:"salt"             env:"PASSWORD_SALT"`
	}
	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"level"   env:"LOG_LEVEL"`
	}

	// MongoDB -.
	MongoDB struct {
		URL      string `env-required:"true" yaml:"url"     env:"DB_URL"`
		Database string `env-required:"true" yaml:"database" env:"DB_DATABASE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
