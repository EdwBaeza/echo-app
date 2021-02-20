package ports

import (
	"github.com/EdwBaeza/echo_app/internal/core/domain"
)

type UsersRepository interface {
	Get(id string) (domain.User, error)
	Save(user domain.User) (domain.User, error)
}
