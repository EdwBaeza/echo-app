package nosql

import (
	"context"
	"sync"
	"time"

	"github.com/EdwinBaeza05/echo_app/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	onceUserRepository     sync.Once
	instanceUserRepository *Repository
)

type Repository struct {
	client  *mongo.Client
	context context.Context
}

//GetClient for mongodb
func (repository *Repository) GetClient() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017/echoapp"))
	repository.client = client
	repository.context, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = repository.client.Connect(repository.context)

	return err
}

//CloseClient for mongodb
func (repository *Repository) CloseClient() {
	repository.client.Disconnect(repository.context)
}

// NewUserRepository return a instance of repository
func NewUserRepository() *Repository {
	onceUserRepository.Do(func() {
		instanceUserRepository = new(Repository)
	})
	return instanceUserRepository
}

//Save user in mongodb
func (repository *Repository) Save(user domain.User) (domain.User, error) {
	collection := repository.client.Database("echoapp").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)

	return user, err
}

//Get user in mongodb
func (repository *Repository) Get(id string) (domain.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	collection := repository.client.Database("echoapp").Collection("users")
	filter := bson.M{"_id": objectID}
	var user domain.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}
