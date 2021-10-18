package ports

import (
	"github.com/EdwBaeza/echo-app/internal/core/domain"
)

type UsersService interface {
	Find(id string) (*domain.User, error)
	All(pageSize int, pageNumber int) (*domain.UserPage, error)
	Create(user domain.User) (*domain.User, error)
}
