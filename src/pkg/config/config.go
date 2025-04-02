package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

type EnvMode string

const (
	EnvModeDev  EnvMode = "DEV"
	EnvModeProd EnvMode = "PROD"
)

type DbConfig struct {
	Host     string `env:"HOST" validate:"required"`
	Port     int    `env:"PORT" validate:"required,gte=0"`
	Name     string `env:"NAME" validate:"required"`
	User     string `env:"USER" validate:"required"`
	Password string `env:"PASSWORD" validate:"required"`
}

type AppConfig struct {
	Host string   `env:"HOST" validate:"required"`
	Port int      `env:"PORT" validate:"required,gte=0"`
	Env  EnvMode  `env:"ENV" validate:"required" envDefault:"PROD"`
	Db   DbConfig `envPrefix:"DB_"`
}

func New() (*AppConfig, error) {
	conf := &AppConfig{}
	err := env.Parse(conf)
	return conf, err
}

func (c *AppConfig) Validate() validator.ValidationErrors {
	return validate.Struct(*c).(validator.ValidationErrors)
}
