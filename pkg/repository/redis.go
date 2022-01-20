package repository

import (
	"github.com/Selahattinn/go-redis/pkg/repository/key"
	"github.com/go-redis/redis"
)

// Redis Repository defines the Redis implementation of Repository interface
type RedisRepository struct {
	cfg     *RedisConfig
	keyRepo key.Repository
}

// RedisConfig defines the Redis Repository configuration
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// RedisConn create a redis client
func RedisConn(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
	})
	return client
}

// NewRedisRepository creates a new Redis Repository
func NewRedisRepository(cfg *RedisConfig) (*RedisRepository, error) {
	client := RedisConn(cfg)
	keyRepo, err := key.NewRedisRepository(client)
	if err != nil {
		return nil, err
	}
	return &RedisRepository{
		cfg:     cfg,
		keyRepo: keyRepo,
	}, nil
}

func (r RedisRepository) GetKeyRepository() key.Repository {
	return r.keyRepo
}
