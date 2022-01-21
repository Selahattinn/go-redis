package key

import (
	"errors"

	"github.com/Selahattinn/go-redis/pkg/model"
	"github.com/Selahattinn/go-redis/pkg/repository"
)

var (
	ErrKeyNotFound = errors.New("key not exist")
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) Get(key string) (string, error) {
	val, err := s.repository.GetKeyRepository().Get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (s *Service) CheckExist(key string) (bool, error) {
	found, err := s.repository.GetKeyRepository().Exist(key)
	if err != nil {
		return found, err
	}
	return found, nil
}

func (s *Service) Store(key model.Key) error {
	err := s.repository.GetKeyRepository().Store(key)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(key model.Key) error {
	found, err := s.repository.GetKeyRepository().Exist(key.ID)
	if err != nil {
		return err
	}
	if found {
		err = s.repository.GetKeyRepository().Store(key)
		if err != nil {
			return err
		}
		return nil
	}

	return ErrKeyNotFound
}

func (s *Service) Delete(key string) error {
	err := s.repository.GetKeyRepository().Delete(key)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() ([]model.Key, error) {
	keys, err := s.repository.GetKeyRepository().GetAll()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (s *Service) DeleteAll() error {
	err := s.repository.GetKeyRepository().DeleteAll()
	if err != nil {
		return err
	}
	return nil
}
