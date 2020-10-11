package ports

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
)

type UsersService interface {
	Get(id int) (domain.User, error)
	Create(user domain.User) error
}