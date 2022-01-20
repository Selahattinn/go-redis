package service

import (
	"github.com/Selahattinn/go-redis/pkg/repository"
	"github.com/Selahattinn/go-redis/pkg/service/key"
)

type Provider struct {
	cfg        *Config
	repository repository.Repository
	keyService *key.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {
	keyService, err := key.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:        cfg,
		repository: repo,
		keyService: keyService,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) GetKeyService() *key.Service {
	return p.keyService
}
