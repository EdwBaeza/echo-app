package ports

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
)

type UsersRepository interface {
	Get(id string) (domain.User, error)
	Save(user domain.User) error
}
