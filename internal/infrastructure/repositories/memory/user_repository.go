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
	data   map[int]domain.User
	lastID int
}

// NewUserRepository return a instance of repository
func NewUserRepository() *Repository {
	onceUserRepository.Do(func() {
		instancePersonRepository = &Repository{
			data:   make(map[int]domain.User),
			lastID: 0,
		}
	})
	return instancePersonRepository
}

// Get user with repository
func (repository *Repository) Get(id int) (domain.User, error) {
	var user domain.User
	user = repository.data[id]
	return user, nil
}

// Create user with repository params
func (repository *Repository) Create(user domain.User) error {
	repository.data[repository.lastID] = user
	repository.lastID++
	return nil
}
