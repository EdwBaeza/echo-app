package memory

import (
	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
	"sync"
)

var(
	onceUserRepository sync.Once
	instancePersonRepository *repository
)

type repository struct{
	data map[int]domain.User
	lastId int
}

func NewUserRepository() *repository{
	onceUserRepository.Do(func(){
		instancePersonRepository = &repository{
			data: make(map[int]domain.User),
			lastId: 0,
		}
	})
	return instancePersonRepository
}

func (repository *repository) Get(id int) (domain.User, error){
	var user domain.User
	user = repository.data[id]
	return user, nil
}

func (repository *repository) Save(user domain.User) error{
	repository.data[repository.lastId] = user
	repository.lastId++
	return nil
}