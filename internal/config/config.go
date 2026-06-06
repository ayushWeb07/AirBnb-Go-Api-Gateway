package config

import "time"

type AppEnvType string

const (
	Development AppEnvType = "development"
	Production  AppEnvType = "production"
)

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	AppEnv       AppEnvType
}
