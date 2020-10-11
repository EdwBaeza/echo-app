package sql

import "github.com/EdwinBaeza05/echo_app/internal/core/domain"

type repository struct{

}

func NewUserRepository() *repository{
	return &repository{}
}

func (repository *repository) Get(id int) (domain.User, error){
	var user domain.User
	return user, nil
}

func (repository *repository) Save(user domain.User) error{
	return nil
}