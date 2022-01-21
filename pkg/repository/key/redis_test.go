package key

import (
	"log"
	"testing"

	"github.com/Selahattinn/go-redis/pkg/model"
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	key = model.Key{
		Value: "val",
		ID:    "1",
	}
	key2 = model.Key{
		Value: "val",
		ID:    "1",
	}
)

// mockInitialize creates a mock and a mock client
func mockInitialize() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	client.FlushAll()
	return client
}

func TestRedisRepository_Get(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Mock Initialize
	mock.Set(key.ID, key.Value, 0)

	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Get function request
	res, err := r.Get(key.ID)
	assert.NoError(t, err)
	assert.Equal(t, key.Value, res)
	client.Close()
}

func TestRedisRepository_Store(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	err = r.Store(key)
	assert.NoError(t, err)

	val, err := mock.Get(key.ID).Result()
	assert.NoError(t, err)
	assert.Equal(t, key.Value, val)
	client.Close()
}

func TestRedisRepository_GetAll(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Set two key to redis
	mock.Set(key.ID, key.Value, 0)
	mock.Set(key2.ID, key2.Value, 0)

	keys, err := r.GetAll()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	assert.Contains(t, keys, key)
	client.Close()

}

func TestRedisRepository_Exist(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// set a key to redis
	mock.Set(key.ID, key.Value, 0)
	found, err := r.Exist(key.ID)
	assert.NoError(t, err)
	assert.Equal(t, found, true)
	client.Close()
}

func TestRedisRepository_ExistFalse(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// set a key to redis
	mock.Set(key.ID, key.Value, 0)
	found, err := r.Exist(key.ID)
	assert.NoError(t, err)
	assert.NotEqual(t, found, false)
	client.Close()
}

func TestRedisRepository_Delete(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// set a key to redis
	mock.Set(key.ID, key.Value, 0)
	err = r.Delete(key.ID)
	assert.NoError(t, err)
	result, err := mock.Get(key.ID).Result()
	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestRedisRepository_DeleteAll(t *testing.T) {
	client := mockInitialize()
	mock := redismock.NewNiceMock(client)
	// Redis Repository Initialize
	r, err := NewRedisRepository(client)
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Set two key to redis
	mock.Set(key.ID, key.Value, 0)
	mock.Set(key2.ID, key2.Value, 0)

	err = r.DeleteAll()
	assert.NoError(t, err)

	result, err := mock.Get(key.ID).Result()
	assert.Error(t, err)
	assert.Equal(t, "", result)

	result, err = mock.Get(key2.ID).Result()
	assert.Error(t, err)
	assert.Equal(t, "", result)
}
