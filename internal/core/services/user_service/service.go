package user_service

import (
	"github.com/EdwBaeza/echo_app/internal/core/domain"
	"github.com/EdwBaeza/echo_app/internal/core/ports"
)

type service struct {
	repository ports.UsersRepository
}

func NewService(repository ports.UsersRepository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) Get(id string) (domain.User, error) {
	return service.repository.Get(id)
}

func (service *service) Create(user domain.User) (domain.User, error) {
	return service.repository.Save(user)
}
