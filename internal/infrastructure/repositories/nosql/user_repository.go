package nosql

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/EdwBaeza/echo-app/internal/core/domain"
	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "users"

var (
	onceUserRepository     sync.Once
	instanceUserRepository *UserRepository
)

type UserRepository struct {
	Repository
}

//BuildClient for mongodb
func (repository *UserRepository) BuildClient() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	repository.client = client
	repository.collection = client.Database(os.Getenv("MONGODB")).Collection(COLLECTION_NAME)
	repository.context, _ = context.WithTimeout(context.Background(), 120*time.Second)
	err = repository.client.Connect(repository.context)

	return err
}

//CloseClient for mongodb
func (repository *UserRepository) CloseClient() {
	repository.client.Disconnect(repository.context)
}

// NewUserRepository return a instance of repository
func NewUserRepository() *UserRepository {
	onceUserRepository.Do(func() {
		instanceUserRepository = new(UserRepository)
	})
	instanceUserRepository.BuildClient()

	return instanceUserRepository
}

//Save user in mongodb
func (repository *UserRepository) Save(user domain.User) (*domain.User, error) {
	result, err := repository.collection.InsertOne(context.TODO(), user)

	if err != nil {
		log.Println("Error insert one: ", err)
		return &user, err
	}
	log.Println("ID created user", result.InsertedID)

	createdUser, _ := repository.Find(result.InsertedID.(primitive.ObjectID).Hex())

	return createdUser, nil
}

//Find user in mongodb
func (repository *UserRepository) Find(id string) (*domain.User, error) {
	user := &domain.User{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error find by id ", err.Error())
		return user, err
	}

	filter := bson.M{"_id": objectID}
	findOneError := repository.collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, findOneError
}

//All user in mongodb
func (repository *UserRepository) All(pageSize int, pageNumber int) (*domain.UserPage, error) {
	count, _ := repository.collection.CountDocuments(repository.context, bson.D{})
	userPage := &domain.UserPage{
		Data: make([]*domain.User, 0),
		Page: domain.Page{
			PageSize:   pageSize,
			PageNumber: pageNumber,
			Count:      int(count),
		},
	}
	paginatedData, err := New(repository.collection).
		Context(repository.context).Limit(int64(pageSize)).
		Page(int64(pageNumber)).
		Aggregate()

	if err != nil {
		log.Println(err)
	}

	for _, raw := range paginatedData.Data {
		user := &domain.User{}
		if marshallErr := bson.Unmarshal(raw, user); marshallErr == nil {
			userPage.Data = append(userPage.Data, user)
		}
	}

	return userPage, err
}
