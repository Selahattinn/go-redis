package key

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) (*RedisRepository, error) {
	return &RedisRepository{
		client: redisClient,
	}, nil
}

func (r *RedisRepository) Get() error {
	status := r.client.Ping()
	fmt.Println(status)
	return nil
}
