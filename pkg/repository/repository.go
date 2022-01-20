package repository

import "github.com/Selahattinn/go-redis/pkg/repository/key"

type Repository interface {
	GetKeyRepository() key.Repository
}
