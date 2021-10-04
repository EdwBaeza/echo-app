package user_service

import (
	"github.com/EdwBaeza/echo-app/internal/core/domain"
	"github.com/EdwBaeza/echo-app/internal/core/ports"
)

type service struct {
	repository ports.UsersRepository
}

func NewService(repository ports.UsersRepository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) Find(id string) (domain.User, error) {
	return service.repository.Find(id)
}

func (service *service) Create(user domain.User) (domain.User, error) {
	return service.repository.Save(user)
}

func (service *service) All() ([]domain.User, error) {
	return service.repository.All()
}
