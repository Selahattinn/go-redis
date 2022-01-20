package service

type Config struct{}

type Service interface {
	GetConfig() *Config
}
