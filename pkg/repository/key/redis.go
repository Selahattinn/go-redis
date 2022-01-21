package key

import (
	"fmt"

	"github.com/Selahattinn/go-redis/pkg/model"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeDuration = 0
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) (*RedisRepository, error) {
	return &RedisRepository{
		client: redisClient,
	}, nil
}

func (r *RedisRepository) Store(key model.Key) error {
	err := r.client.Set(key.ID, key.Value, defaultTimeDuration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisRepository) GetAll() ([]model.Key, error) {
	keys, err := r.client.Keys("*").Result()
	if err != nil {
		return nil, err
	}
	var keyArray []model.Key
	for _, key := range keys {
		var tmpKey model.Key
		tmpKey.ID = key
		val, err := r.client.Get(key).Result()
		if err != nil {
			logrus.WithError(err).Info("Error getting get all Info")
		}
		tmpKey.Value = val
		keyArray = append(keyArray, tmpKey)
	}
	return keyArray, nil
}
func (r *RedisRepository) Exist(key string) (bool, error) {
	val, err := r.client.Exists(key).Result()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if val > 0 {
		return true, nil
	}
	return false, nil
}

func (r *RedisRepository) Delete(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) DeleteAll() error {
	err := r.client.FlushAll().Err()
	if err != nil {
		return err
	}
	return nil
}
