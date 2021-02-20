package ports

import (
	"github.com/EdwBaeza/echo_app/internal/core/domain"
)

type UsersService interface {
	Get(id string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
}
