package repository

import (
	"context"
	"errors"

	"github.com/go-redis/redis"
)

// RedisConfig defines the Redis Repository configuration
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Redis struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.Background()
)

func NewDatabase(cfg *RedisConfig) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Redis{
		Client: client,
	}, nil
}
