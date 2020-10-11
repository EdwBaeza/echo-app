package user

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
	"github.com/EdwinBaeza05/echo_app/internal/core/ports"
)

type service struct{
	repository ports.UsersRepository
}

func NewService(repository ports.UsersRepository) *service{
	return &service{
		repository: repository,
	}
}

func (service *service) Get(id int) (domain.User, error){
	return service.repository.Get(id)
}

func (service *service) Create(user domain.User) error{
	return service.repository.Create(user)
}