package key

import "github.com/Selahattinn/go-redis/pkg/model"

type Reader interface {
	Get(key string) (string, error)
	GetAll() ([]model.Key, error)
	Exist(key string) (bool, error)
}

type Writer interface {
	Store(key model.Key) error
	Delete(key string) error
	DeleteAll() error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
