package memory

import (
	"sync"

	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
)

var (
	onceUserRepository       sync.Once
	instancePersonRepository *Repository
)

//Repository in memory
type Repository struct {
	data   map[string]domain.User
	lastID string
}

// NewUserRepository return a instance of repository
func NewUserRepository() *Repository {
	onceUserRepository.Do(func() {
		instancePersonRepository = &Repository{
			data:   make(map[string]domain.User),
			lastID: "",
		}
	})
	return instancePersonRepository
}

// Get user with repository
func (repository *Repository) Get(id string) (domain.User, error) {
	var user domain.User
	user = repository.data[id]
	return user, nil
}

// Save user with repository params
func (repository *Repository) Save(user domain.User) (domain.User, error) {
	repository.data[user.ID] = user
	repository.lastID = user.ID
	return user, nil
}
