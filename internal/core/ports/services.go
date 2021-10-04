package ports

import (
	"github.com/EdwBaeza/echo-app/internal/core/domain"
)

type UsersService interface {
	Find(id string) (domain.User, error)
	All() ([]domain.User, error)
	Create(user domain.User) (domain.User, error)
}
