package service

import "github.com/Selahattinn/go-redis/pkg/service/key"

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetKeyService() *key.Service
}
