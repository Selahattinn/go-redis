package key

import "github.com/Selahattinn/go-redis/pkg/repository"

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) Get() error {

	return nil
}

func (s *Service) CheckExist() error {
	return nil
}

func (s *Service) Store() error {
	return nil
}

func (s *Service) Update() error {
	return nil
}

func (s *Service) Delete() error {
	return nil
}

func (s *Service) GetAll() error {
	return nil
}

func (s *Service) DeleteAll() error {
	return nil
}
