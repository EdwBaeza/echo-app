package sql

import "github.com/EdwinBaeza05/echo_app/internal/core/domain"

type repository struct {
}

// NewUserRepository should return a user_repository
func NewUserRepository() *repository {
	return &repository{}
}

func (repository *repository) Get(id int) (domain.User, error) {
	var user domain.User
	return user, nil
}

func (repository *repository) Save(user domain.User) (domain.User, error) {
	return user, nil
}
