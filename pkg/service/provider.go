package service

import (
	"github.com/Selahattinn/go-redis/pkg/repository"
)

type Provider struct {
	cfg        *Config
	repository repository.Redis
}

func NewProvider(cfg *Config, repo repository.Redis) (*Provider, error) {

	return &Provider{
		cfg:        cfg,
		repository: repo,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}
